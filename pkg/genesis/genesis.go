package genesis

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/ipfs/go-blockservice"
	"github.com/ipfs/go-cid"
	offline "github.com/ipfs/go-ipfs-exchange-offline"
	cbor "github.com/ipfs/go-ipld-cbor"
	logging "github.com/ipfs/go-log/v2"
	"github.com/ipfs/go-merkledag"
	"github.com/ipld/go-car"
	"github.com/mitchellh/go-homedir"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/venus/pkg/config"
	"github.com/filecoin-project/venus/pkg/constants"
	"github.com/filecoin-project/venus/pkg/gen"
	genesis2 "github.com/filecoin-project/venus/pkg/gen/genesis"
	"github.com/filecoin-project/venus/pkg/repo"
	"github.com/filecoin-project/venus/pkg/state/tree"
	"github.com/filecoin-project/venus/pkg/util/blockstoreutil"
	"github.com/filecoin-project/venus/pkg/vm"
	"github.com/filecoin-project/venus/venus-shared/types"

	"github.com/filecoin-project/venus/fixtures/assets"
	"github.com/filecoin-project/venus/fixtures/networks"
)

var glog = logging.Logger("genesis")

// InitFunc is the signature for function that is used to create a genesis block.
type InitFunc func(cst cbor.IpldStore, bs blockstoreutil.Blockstore) (*types.BlockHeader, error)

// Ticket is the ticket to place in the genesis block header (which can't be derived from a prior ticket),
// used in the evaluation of the messages in the genesis block,
// and *also* the ticket value used when computing the genesis state (the parent state of the genesis block).
var Ticket = types.Ticket{
	VRFProof: []byte{
		0xec, 0xec, 0xec, 0xec, 0xec, 0xec, 0xec, 0xec, 0xec, 0xec, 0xec, 0xec, 0xec, 0xec, 0xec, 0xec,
		0xec, 0xec, 0xec, 0xec, 0xec, 0xec, 0xec, 0xec, 0xec, 0xec, 0xec, 0xec, 0xec, 0xec, 0xec, 0xec,
	},
}

// VM is the view into the LegacyVM used during genesis block creation.
type VM interface {
	ApplyGenesisMessage(from address.Address, to address.Address, method abi.MethodNum, value abi.TokenAmount, params interface{}) (*vm.Ret, error)
	Flush(ctx context.Context) (tree.Root, error)
}

//MakeGenesis return a func to construct a genesis block
func MakeGenesis(ctx context.Context, rep repo.Repo, outFile, genesisTemplate string, para *config.ForkUpgradeConfig) InitFunc {
	return func(_ cbor.IpldStore, bs blockstoreutil.Blockstore) (*types.BlockHeader, error) {
		glog.Warn("Generating new random genesis block, note that this SHOULD NOT happen unless you are setting up new network")
		genesisTemplate, err := homedir.Expand(genesisTemplate)
		if err != nil {
			return nil, err
		}

		fdata, err := ioutil.ReadFile(genesisTemplate)
		if err != nil {
			return nil, fmt.Errorf("reading preseals json: %w", err)
		}

		var template genesis2.Template
		if err := json.Unmarshal(fdata, &template); err != nil {
			return nil, err
		}

		if template.Timestamp == 0 {
			template.Timestamp = uint64(constants.Clock.Now().Unix())
		}

		// TODO potentially replace this cached blockstore by a CBOR cache.
		cbs, err := blockstoreutil.CachedBlockstore(ctx, bs, blockstoreutil.DefaultCacheOpts())
		if err != nil {
			return nil, err
		}

		b, err := genesis2.MakeGenesisBlock(context.TODO(), rep, cbs, template, para)
		if err != nil {
			return nil, fmt.Errorf("make genesis block: %w", err)
		}

		fmt.Printf("GENESIS MINER ADDRESS: t0%d\n", genesis2.MinerStart)

		f, err := os.OpenFile(outFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			return nil, err
		}

		offl := offline.Exchange(cbs)
		blkserv := blockservice.New(cbs, offl)
		dserv := merkledag.NewDAGService(blkserv)

		if err := car.WriteCarWithWalker(context.TODO(), dserv, []cid.Cid{b.Genesis.Cid()}, f, gen.CarWalkFunc); err != nil {
			return nil, err
		}

		glog.Infof("WRITING GENESIS FILE AT %s", f.Name())

		if err := f.Close(); err != nil {
			return nil, err
		}

		return b.Genesis, nil

	}
}

func LoadGenesis(ctx context.Context, rep repo.Repo, sourceName string, network string) (InitFunc, error) {
	var (
		source io.ReadCloser
		err    error
	)

	if sourceName == "" {
		networkType, err := networks.GetNetworkFromName(network)
		if err != nil {
			return nil, err
		}

		bs, err := assets.GetGenesis(networkType)
		if err != nil {
			return nil, err
		}
		source = ioutil.NopCloser(bytes.NewReader(bs))
	} else {
		source, err = openGenesisSource(sourceName)
		if err != nil {
			return nil, err
		}
	}

	defer func() { _ = source.Close() }()

	genesisBlk, err := extractGenesisBlock(ctx, source, rep)
	if err != nil {
		return nil, err
	}

	gif := func(cst cbor.IpldStore, bs blockstoreutil.Blockstore) (*types.BlockHeader, error) {
		return genesisBlk, err
	}

	return gif, nil
}

func extractGenesisBlock(ctx context.Context, source io.ReadCloser, rep repo.Repo) (*types.BlockHeader, error) {
	bs := rep.Datastore()
	ch, err := car.LoadCar(ctx, bs, source)
	if err != nil {
		return nil, err
	}

	// need to check if we are being handed a car file with a single genesis block or an entire chain.
	bsBlk, err := bs.Get(ctx, ch.Roots[0])
	if err != nil {
		return nil, err
	}
	cur, err := types.DecodeBlock(bsBlk.RawData())
	if err != nil {
		return nil, err
	}

	return cur, nil
}

func openGenesisSource(sourceName string) (io.ReadCloser, error) {
	sourceURL, err := url.Parse(sourceName)
	if err != nil {
		return nil, fmt.Errorf("invalid filepath or URL for genesis file: %s", sourceURL)
	}
	var source io.ReadCloser
	if sourceURL.Scheme == "http" || sourceURL.Scheme == "https" {
		// NOTE: This code is temporary. It allows downloading a genesis block via HTTP(S) to be able to join a
		// recently deployed staging devnet.
		response, err := http.Get(sourceName)
		if err != nil {
			return nil, err
		}
		source = response.Body
	} else if sourceURL.Scheme != "" {
		return nil, fmt.Errorf("unsupported protocol for genesis file: %s", sourceURL.Scheme)
	} else {
		file, err := os.Open(sourceName)
		if err != nil {
			return nil, err
		}
		source = file
	}
	return source, nil
}
