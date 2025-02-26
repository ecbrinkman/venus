package testhelpers

import (
	"context"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"testing"
	"time"

	"github.com/filecoin-project/venus/pkg/net/exchange"
	"github.com/filecoin-project/venus/venus-shared/types"
	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p-core/connmgr"
	"github.com/libp2p/go-libp2p-core/event"
	"github.com/libp2p/go-libp2p-core/host"
	inet "github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/peerstore"
	"github.com/libp2p/go-libp2p-core/protocol"
	ma "github.com/multiformats/go-multiaddr"
	mh "github.com/multiformats/go-multihash"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

var _ host.Host = &FakeHost{}
var _ connmgr.ConnManager = &FakeCMgr{}

// FakeHost is a test host.Host
type FakeHost struct {
	ConnectImpl func(context.Context, peer.AddrInfo) error
}

// NewFakeHost constructs a FakeHost with no other parameters needed
func NewFakeHost() host.Host {
	nopfunc := func(_ context.Context, _ peer.AddrInfo) error { return nil }
	return &FakeHost{
		ConnectImpl: nopfunc,
	}
}

// minimal implementation of host.Host interface

func (fh *FakeHost) Addrs() []ma.Multiaddr { panic("not implemented") } // nolint: golint
func (fh *FakeHost) Close() error          { panic("not implemented") } // nolint: golint
func (fh *FakeHost) ConnManager() connmgr.ConnManager {
	return &FakeCMgr{}
} // nolint: golint
func (fh *FakeHost) Connect(ctx context.Context, pi peer.AddrInfo) error { // nolint: golint
	return fh.ConnectImpl(ctx, pi)
}
func (fh *FakeHost) EventBus() event.Bus                              { panic("not implemented") } //nolint: golint
func (fh *FakeHost) ID() peer.ID                                      { panic("not implemented") } // nolint: golint
func (fh *FakeHost) Network() inet.Network                            { panic("not implemented") } // nolint: golint
func (fh *FakeHost) Mux() protocol.Switch                             { panic("not implemented") } // nolint: golint
func (fh *FakeHost) Peerstore() peerstore.Peerstore                   { panic("not implemented") } // nolint: golint
func (fh *FakeHost) RemoveStreamHandler(protocol.ID)                  { panic("not implemented") } // nolint: golint
func (fh *FakeHost) SetStreamHandler(protocol.ID, inet.StreamHandler) { panic("not implemented") } // nolint: golint
func (fh *FakeHost) SetStreamHandlerMatch(protocol.ID, func(string) bool, inet.StreamHandler) { // nolint: golint
	panic("not implemented")
}

type FakeCMgr struct {
}

func (f FakeCMgr) TagPeer(id peer.ID, s string, i int) {}

func (f FakeCMgr) UntagPeer(p peer.ID, tag string) {}

func (f FakeCMgr) UpsertTag(p peer.ID, tag string, upsert func(int) int) {}

func (f FakeCMgr) GetTagInfo(p peer.ID) *connmgr.TagInfo {
	panic("implement me")
}

func (f FakeCMgr) TrimOpenConns(ctx context.Context) {
	panic("implement me")
}

func (f FakeCMgr) Notifee() inet.Notifiee {
	panic("implement me")
}

func (f FakeCMgr) Protect(id peer.ID, tag string) {
	panic("implement me")
}

func (f FakeCMgr) Unprotect(id peer.ID, tag string) (protected bool) {
	panic("implement me")
}

func (f FakeCMgr) IsProtected(id peer.ID, tag string) (protected bool) {
	panic("implement me")
}

func (f FakeCMgr) Close() error {
	panic("implement me")
}

// NewStream is required for the host.Host interface; returns a new FakeStream.
func (fh *FakeHost) NewStream(context.Context, peer.ID, ...protocol.ID) (inet.Stream, error) { // nolint: golint
	return newFakeStream(), nil
}

var _ inet.Dialer = &FakeDialer{}

// FakeDialer is a test inet.Dialer
type FakeDialer struct {
	PeersImpl func() []peer.ID
}

// Minimal implementation of the inet.Dialer interface

// Peers returns a fake inet.Dialer PeersImpl
func (fd *FakeDialer) Peers() []peer.ID {
	return fd.PeersImpl()
}
func (fd *FakeDialer) Peerstore() peerstore.Peerstore                       { panic("not implemented") }
func (fd *FakeDialer) LocalPeer() peer.ID                                   { panic("not implemented") }
func (fd *FakeDialer) DialPeer(context.Context, peer.ID) (inet.Conn, error) { panic("not implemented") }
func (fd *FakeDialer) ClosePeer(peer.ID) error                              { panic("not implemented") }
func (fd *FakeDialer) Connectedness(peer.ID) inet.Connectedness             { panic("not implemented") }
func (fd *FakeDialer) Conns() []inet.Conn                                   { panic("not implemented") }
func (fd *FakeDialer) ConnsToPeer(peer.ID) []inet.Conn                      { panic("not implemented") }
func (fd *FakeDialer) Notify(inet.Notifiee)                                 { panic("not implemented") }
func (fd *FakeDialer) StopNotify(inet.Notifiee)                             { panic("not implemented") }

// fakeStream is a test inet.Stream
type fakeStream struct {
	_   inet.MuxedStream
	pid protocol.ID
}

var _ inet.Stream = &fakeStream{}

func newFakeStream() *fakeStream { return &fakeStream{} }

// Minimal implementation of the inet.Stream interface
func (fs *fakeStream) ID() string                         { return "" }
func (fs *fakeStream) Protocol() protocol.ID              { return fs.pid }            // nolint: golint
func (fs *fakeStream) SetProtocol(id protocol.ID) error   { fs.pid = id; return nil }  // nolint: golint
func (fs *fakeStream) Stat() inet.Stats                   { panic("not implemented") } // nolint: golint
func (fs *fakeStream) Conn() inet.Conn                    { panic("not implemented") } // nolint: golint
func (fs *fakeStream) Write(_ []byte) (int, error)        { return 1, nil }            // nolint: golint
func (fs *fakeStream) Read(_ []byte) (int, error)         { return 1, nil }            // nolint: golint
func (fs *fakeStream) Close() error                       { return nil }               // nolint: golint
func (fs *fakeStream) Reset() error                       { return nil }               // nolint: golint
func (fs *fakeStream) SetDeadline(_ time.Time) error      { return nil }               // nolint: golint
func (fs *fakeStream) SetReadDeadline(_ time.Time) error  { return nil }               // nolint: golint
func (fs *fakeStream) SetWriteDeadline(_ time.Time) error { return nil }               // nolint: golint
func (fs *fakeStream) CloseWrite() error                  { panic("implement me") }
func (fs *fakeStream) CloseRead() error                   { panic("implement me") }
func (fs *fakeStream) Scope() inet.StreamScope            { panic("implement me") }

// RandPeerID is a libp2p random peer ID generator.
// These peer.ID generators were copied from libp2p/go-testutil. We didn't bring in the
// whole repo as a dependency because we only need this small bit. However if we find
// ourselves using more and more pieces we should just take a dependency on it.
func RandPeerID() (peer.ID, error) {
	buf := make([]byte, 16)
	if n, err := rand.Read(buf); n != 16 || err != nil {
		if n != 16 && err == nil {
			err = errors.New("couldnt read 16 random bytes")
		}
		panic(err)
	}
	h, _ := mh.Sum(buf, mh.SHA2_256, -1)
	return peer.ID(h), nil
}

// RequireIntPeerID takes in an integer and creates a unique peer id for it.
func RequireIntPeerID(t *testing.T, i int64) peer.ID {
	buf := make([]byte, 16)
	n := binary.PutVarint(buf, i)
	h, err := mh.Sum(buf[:n], mh.IDENTITY, -1)
	require.NoError(t, err)
	pid, err := peer.IDFromBytes(h)
	require.NoError(t, err)
	return pid
}

// TestFetcher is an object with the same method set as Fetcher plus a method
// for adding blocks to the source.  It is used to implement an object that
// behaves like Fetcher but does not go to the network for use in tests.
type TestFetcher struct {
	sourceBlocks map[string]*types.BlockHeader // sourceBlocks maps block cid strings to blocks.
}

// NewTestFetcher returns a TestFetcher with no source blocks.
func NewTestFetcher() *TestFetcher {
	return &TestFetcher{
		sourceBlocks: make(map[string]*types.BlockHeader),
	}
}

// AddSourceBlocks adds the input blocks to the fetcher source.
func (f *TestFetcher) AddSourceBlocks(blocks ...*types.BlockHeader) {
	for _, block := range blocks {
		f.sourceBlocks[block.Cid().String()] = block
	}
}

// FetchTipSets fetchs the tipset at `tsKey` from the network using the fetchers `sourceBlocks`.
func (f *TestFetcher) FetchTipSets(ctx context.Context, tsKey types.TipSetKey, from peer.ID, done func(t *types.TipSet) (bool, error)) ([]*types.TipSet, error) {
	var out []*types.TipSet
	cur := tsKey
	for {
		res, err := f.GetBlocks(ctx, cur.Cids())
		if err != nil {
			return nil, err
		}

		ts, err := types.NewTipSet(res)
		if err != nil {
			return nil, err
		}

		out = append(out, ts)
		ok, err := done(ts)
		if err != nil {
			return nil, err
		}
		if ok {
			break
		}

		cur = ts.Parents()
	}

	return out, nil
}

// FetchTipSetHeaders fetches the tipset at `tsKey` but not messages
func (f *TestFetcher) FetchTipSetHeaders(ctx context.Context, tsKey types.TipSetKey, from peer.ID, done func(t *types.TipSet) (bool, error)) ([]*types.TipSet, error) {
	return f.FetchTipSets(ctx, tsKey, from, done)
}

// GetBlocks returns any blocks in the source with matching cids.
func (f *TestFetcher) GetBlocks(ctx context.Context, cids []cid.Cid) ([]*types.BlockHeader, error) {
	var ret []*types.BlockHeader
	for _, c := range cids {
		if block, ok := f.sourceBlocks[c.String()]; ok {
			ret = append(ret, block)
		} else {
			return nil, fmt.Errorf("failed to fetch block: %s", c.String())
		}
	}
	return ret, nil
}

func NewTestExchange() *TestExchange {
	return &TestExchange{
		sourceBlocks: make(map[string]*types.BlockHeader),
	}
}

// AddSourceBlocks adds the input blocks to the fetcher source.
func (f *TestExchange) AddSourceBlocks(blocks ...*types.BlockHeader) {
	for _, block := range blocks {
		f.sourceBlocks[block.Cid().String()] = block
	}
}

type TestExchange struct {
	sourceBlocks map[string]*types.BlockHeader // s
}

func (f *TestExchange) GetBlocks(ctx context.Context, tsk types.TipSetKey, count int) ([]*types.TipSet, error) {
	panic("implement me")
}

func (f *TestExchange) GetChainMessages(ctx context.Context, tipsets []*types.TipSet) ([]*exchange.CompactedMessages, error) {
	panic("implement me")
}

func (f *TestExchange) GetFullTipSet(ctx context.Context, peer []peer.ID, tsk types.TipSetKey) (*types.FullTipSet, error) {
	panic("implement me")
}

func (f *TestExchange) AddPeer(peer peer.ID) {}

func (f *TestExchange) RemovePeer(peer peer.ID) {}
