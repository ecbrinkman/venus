// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/filecoin-project/venus/venus-shared/api/messager (interfaces: IMessager)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"
	time "time"

	address "github.com/filecoin-project/go-address"
	internal "github.com/filecoin-project/venus/venus-shared/internal"
	types "github.com/filecoin-project/venus/venus-shared/types"
	messager "github.com/filecoin-project/venus/venus-shared/types/messager"
	gomock "github.com/golang/mock/gomock"
	cid "github.com/ipfs/go-cid"
	peer "github.com/libp2p/go-libp2p-core/peer"
)

// MockIMessager is a mock of IMessager interface.
type MockIMessager struct {
	ctrl     *gomock.Controller
	recorder *MockIMessagerMockRecorder
}

// MockIMessagerMockRecorder is the mock recorder for MockIMessager.
type MockIMessagerMockRecorder struct {
	mock *MockIMessager
}

// NewMockIMessager creates a new mock instance.
func NewMockIMessager(ctrl *gomock.Controller) *MockIMessager {
	mock := &MockIMessager{ctrl: ctrl}
	mock.recorder = &MockIMessagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIMessager) EXPECT() *MockIMessagerMockRecorder {
	return m.recorder
}

// ActiveAddress mocks base method.
func (m *MockIMessager) ActiveAddress(arg0 context.Context, arg1 address.Address) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActiveAddress", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ActiveAddress indicates an expected call of ActiveAddress.
func (mr *MockIMessagerMockRecorder) ActiveAddress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActiveAddress", reflect.TypeOf((*MockIMessager)(nil).ActiveAddress), arg0, arg1)
}

// ClearUnFillMessage mocks base method.
func (m *MockIMessager) ClearUnFillMessage(arg0 context.Context, arg1 address.Address) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClearUnFillMessage", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClearUnFillMessage indicates an expected call of ClearUnFillMessage.
func (mr *MockIMessagerMockRecorder) ClearUnFillMessage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearUnFillMessage", reflect.TypeOf((*MockIMessager)(nil).ClearUnFillMessage), arg0, arg1)
}

// DeleteAddress mocks base method.
func (m *MockIMessager) DeleteAddress(arg0 context.Context, arg1 address.Address) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAddress", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAddress indicates an expected call of DeleteAddress.
func (mr *MockIMessagerMockRecorder) DeleteAddress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAddress", reflect.TypeOf((*MockIMessager)(nil).DeleteAddress), arg0, arg1)
}

// DeleteNode mocks base method.
func (m *MockIMessager) DeleteNode(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNode", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteNode indicates an expected call of DeleteNode.
func (mr *MockIMessagerMockRecorder) DeleteNode(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNode", reflect.TypeOf((*MockIMessager)(nil).DeleteNode), arg0, arg1)
}

// ForbiddenAddress mocks base method.
func (m *MockIMessager) ForbiddenAddress(arg0 context.Context, arg1 address.Address) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForbiddenAddress", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ForbiddenAddress indicates an expected call of ForbiddenAddress.
func (mr *MockIMessagerMockRecorder) ForbiddenAddress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForbiddenAddress", reflect.TypeOf((*MockIMessager)(nil).ForbiddenAddress), arg0, arg1)
}

// ForcePushMessage mocks base method.
func (m *MockIMessager) ForcePushMessage(arg0 context.Context, arg1 string, arg2 *internal.Message, arg3 *messager.SendSpec) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForcePushMessage", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ForcePushMessage indicates an expected call of ForcePushMessage.
func (mr *MockIMessagerMockRecorder) ForcePushMessage(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForcePushMessage", reflect.TypeOf((*MockIMessager)(nil).ForcePushMessage), arg0, arg1, arg2, arg3)
}

// ForcePushMessageWithId mocks base method.
func (m *MockIMessager) ForcePushMessageWithId(arg0 context.Context, arg1, arg2 string, arg3 *internal.Message, arg4 *messager.SendSpec) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForcePushMessageWithId", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ForcePushMessageWithId indicates an expected call of ForcePushMessageWithId.
func (mr *MockIMessagerMockRecorder) ForcePushMessageWithId(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForcePushMessageWithId", reflect.TypeOf((*MockIMessager)(nil).ForcePushMessageWithId), arg0, arg1, arg2, arg3, arg4)
}

// GetAddress mocks base method.
func (m *MockIMessager) GetAddress(arg0 context.Context, arg1 address.Address) (*messager.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAddress", arg0, arg1)
	ret0, _ := ret[0].(*messager.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAddress indicates an expected call of GetAddress.
func (mr *MockIMessagerMockRecorder) GetAddress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAddress", reflect.TypeOf((*MockIMessager)(nil).GetAddress), arg0, arg1)
}

// GetMessageByFromAndNonce mocks base method.
func (m *MockIMessager) GetMessageByFromAndNonce(arg0 context.Context, arg1 address.Address, arg2 uint64) (*messager.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMessageByFromAndNonce", arg0, arg1, arg2)
	ret0, _ := ret[0].(*messager.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMessageByFromAndNonce indicates an expected call of GetMessageByFromAndNonce.
func (mr *MockIMessagerMockRecorder) GetMessageByFromAndNonce(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMessageByFromAndNonce", reflect.TypeOf((*MockIMessager)(nil).GetMessageByFromAndNonce), arg0, arg1, arg2)
}

// GetMessageBySignedCid mocks base method.
func (m *MockIMessager) GetMessageBySignedCid(arg0 context.Context, arg1 cid.Cid) (*messager.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMessageBySignedCid", arg0, arg1)
	ret0, _ := ret[0].(*messager.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMessageBySignedCid indicates an expected call of GetMessageBySignedCid.
func (mr *MockIMessagerMockRecorder) GetMessageBySignedCid(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMessageBySignedCid", reflect.TypeOf((*MockIMessager)(nil).GetMessageBySignedCid), arg0, arg1)
}

// GetMessageByUid mocks base method.
func (m *MockIMessager) GetMessageByUid(arg0 context.Context, arg1 string) (*messager.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMessageByUid", arg0, arg1)
	ret0, _ := ret[0].(*messager.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMessageByUid indicates an expected call of GetMessageByUid.
func (mr *MockIMessagerMockRecorder) GetMessageByUid(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMessageByUid", reflect.TypeOf((*MockIMessager)(nil).GetMessageByUid), arg0, arg1)
}

// GetMessageByUnsignedCid mocks base method.
func (m *MockIMessager) GetMessageByUnsignedCid(arg0 context.Context, arg1 cid.Cid) (*messager.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMessageByUnsignedCid", arg0, arg1)
	ret0, _ := ret[0].(*messager.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMessageByUnsignedCid indicates an expected call of GetMessageByUnsignedCid.
func (mr *MockIMessagerMockRecorder) GetMessageByUnsignedCid(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMessageByUnsignedCid", reflect.TypeOf((*MockIMessager)(nil).GetMessageByUnsignedCid), arg0, arg1)
}

// GetNode mocks base method.
func (m *MockIMessager) GetNode(arg0 context.Context, arg1 string) (*messager.Node, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNode", arg0, arg1)
	ret0, _ := ret[0].(*messager.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNode indicates an expected call of GetNode.
func (mr *MockIMessagerMockRecorder) GetNode(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNode", reflect.TypeOf((*MockIMessager)(nil).GetNode), arg0, arg1)
}

// GetSharedParams mocks base method.
func (m *MockIMessager) GetSharedParams(arg0 context.Context) (*messager.SharedSpec, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSharedParams", arg0)
	ret0, _ := ret[0].(*messager.SharedSpec)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSharedParams indicates an expected call of GetSharedParams.
func (mr *MockIMessagerMockRecorder) GetSharedParams(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSharedParams", reflect.TypeOf((*MockIMessager)(nil).GetSharedParams), arg0)
}

// HasAddress mocks base method.
func (m *MockIMessager) HasAddress(arg0 context.Context, arg1 address.Address) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasAddress", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasAddress indicates an expected call of HasAddress.
func (mr *MockIMessagerMockRecorder) HasAddress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasAddress", reflect.TypeOf((*MockIMessager)(nil).HasAddress), arg0, arg1)
}

// HasMessageByUid mocks base method.
func (m *MockIMessager) HasMessageByUid(arg0 context.Context, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasMessageByUid", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasMessageByUid indicates an expected call of HasMessageByUid.
func (mr *MockIMessagerMockRecorder) HasMessageByUid(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasMessageByUid", reflect.TypeOf((*MockIMessager)(nil).HasMessageByUid), arg0, arg1)
}

// HasNode mocks base method.
func (m *MockIMessager) HasNode(arg0 context.Context, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasNode", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasNode indicates an expected call of HasNode.
func (mr *MockIMessagerMockRecorder) HasNode(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasNode", reflect.TypeOf((*MockIMessager)(nil).HasNode), arg0, arg1)
}

// ListAddress mocks base method.
func (m *MockIMessager) ListAddress(arg0 context.Context) ([]*messager.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAddress", arg0)
	ret0, _ := ret[0].([]*messager.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAddress indicates an expected call of ListAddress.
func (mr *MockIMessagerMockRecorder) ListAddress(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAddress", reflect.TypeOf((*MockIMessager)(nil).ListAddress), arg0)
}

// ListBlockedMessage mocks base method.
func (m *MockIMessager) ListBlockedMessage(arg0 context.Context, arg1 address.Address, arg2 time.Duration) ([]*messager.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBlockedMessage", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*messager.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBlockedMessage indicates an expected call of ListBlockedMessage.
func (mr *MockIMessagerMockRecorder) ListBlockedMessage(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBlockedMessage", reflect.TypeOf((*MockIMessager)(nil).ListBlockedMessage), arg0, arg1, arg2)
}

// ListFailedMessage mocks base method.
func (m *MockIMessager) ListFailedMessage(arg0 context.Context) ([]*messager.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFailedMessage", arg0)
	ret0, _ := ret[0].([]*messager.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFailedMessage indicates an expected call of ListFailedMessage.
func (mr *MockIMessagerMockRecorder) ListFailedMessage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFailedMessage", reflect.TypeOf((*MockIMessager)(nil).ListFailedMessage), arg0)
}

// ListMessage mocks base method.
func (m *MockIMessager) ListMessage(arg0 context.Context) ([]*messager.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMessage", arg0)
	ret0, _ := ret[0].([]*messager.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMessage indicates an expected call of ListMessage.
func (mr *MockIMessagerMockRecorder) ListMessage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMessage", reflect.TypeOf((*MockIMessager)(nil).ListMessage), arg0)
}

// ListMessageByAddress mocks base method.
func (m *MockIMessager) ListMessageByAddress(arg0 context.Context, arg1 address.Address) ([]*messager.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMessageByAddress", arg0, arg1)
	ret0, _ := ret[0].([]*messager.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMessageByAddress indicates an expected call of ListMessageByAddress.
func (mr *MockIMessagerMockRecorder) ListMessageByAddress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMessageByAddress", reflect.TypeOf((*MockIMessager)(nil).ListMessageByAddress), arg0, arg1)
}

// ListMessageByFromState mocks base method.
func (m *MockIMessager) ListMessageByFromState(arg0 context.Context, arg1 address.Address, arg2 messager.MessageState, arg3 bool, arg4, arg5 int) ([]*messager.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMessageByFromState", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].([]*messager.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMessageByFromState indicates an expected call of ListMessageByFromState.
func (mr *MockIMessagerMockRecorder) ListMessageByFromState(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMessageByFromState", reflect.TypeOf((*MockIMessager)(nil).ListMessageByFromState), arg0, arg1, arg2, arg3, arg4, arg5)
}

// ListNode mocks base method.
func (m *MockIMessager) ListNode(arg0 context.Context) ([]*messager.Node, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListNode", arg0)
	ret0, _ := ret[0].([]*messager.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListNode indicates an expected call of ListNode.
func (mr *MockIMessagerMockRecorder) ListNode(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNode", reflect.TypeOf((*MockIMessager)(nil).ListNode), arg0)
}

// MarkBadMessage mocks base method.
func (m *MockIMessager) MarkBadMessage(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarkBadMessage", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// MarkBadMessage indicates an expected call of MarkBadMessage.
func (mr *MockIMessagerMockRecorder) MarkBadMessage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkBadMessage", reflect.TypeOf((*MockIMessager)(nil).MarkBadMessage), arg0, arg1)
}

// NetAddrsListen mocks base method.
func (m *MockIMessager) NetAddrsListen(arg0 context.Context) (peer.AddrInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NetAddrsListen", arg0)
	ret0, _ := ret[0].(peer.AddrInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NetAddrsListen indicates an expected call of NetAddrsListen.
func (mr *MockIMessagerMockRecorder) NetAddrsListen(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetAddrsListen", reflect.TypeOf((*MockIMessager)(nil).NetAddrsListen), arg0)
}

// NetConnect mocks base method.
func (m *MockIMessager) NetConnect(arg0 context.Context, arg1 peer.AddrInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NetConnect", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// NetConnect indicates an expected call of NetConnect.
func (mr *MockIMessagerMockRecorder) NetConnect(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetConnect", reflect.TypeOf((*MockIMessager)(nil).NetConnect), arg0, arg1)
}

// NetFindPeer mocks base method.
func (m *MockIMessager) NetFindPeer(arg0 context.Context, arg1 peer.ID) (peer.AddrInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NetFindPeer", arg0, arg1)
	ret0, _ := ret[0].(peer.AddrInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NetFindPeer indicates an expected call of NetFindPeer.
func (mr *MockIMessagerMockRecorder) NetFindPeer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetFindPeer", reflect.TypeOf((*MockIMessager)(nil).NetFindPeer), arg0, arg1)
}

// NetPeers mocks base method.
func (m *MockIMessager) NetPeers(arg0 context.Context) ([]peer.AddrInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NetPeers", arg0)
	ret0, _ := ret[0].([]peer.AddrInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NetPeers indicates an expected call of NetPeers.
func (mr *MockIMessagerMockRecorder) NetPeers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetPeers", reflect.TypeOf((*MockIMessager)(nil).NetPeers), arg0)
}

// PushMessage mocks base method.
func (m *MockIMessager) PushMessage(arg0 context.Context, arg1 *internal.Message, arg2 *messager.SendSpec) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PushMessage", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PushMessage indicates an expected call of PushMessage.
func (mr *MockIMessagerMockRecorder) PushMessage(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PushMessage", reflect.TypeOf((*MockIMessager)(nil).PushMessage), arg0, arg1, arg2)
}

// PushMessageWithId mocks base method.
func (m *MockIMessager) PushMessageWithId(arg0 context.Context, arg1 string, arg2 *internal.Message, arg3 *messager.SendSpec) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PushMessageWithId", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PushMessageWithId indicates an expected call of PushMessageWithId.
func (mr *MockIMessagerMockRecorder) PushMessageWithId(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PushMessageWithId", reflect.TypeOf((*MockIMessager)(nil).PushMessageWithId), arg0, arg1, arg2, arg3)
}

// RecoverFailedMsg mocks base method.
func (m *MockIMessager) RecoverFailedMsg(arg0 context.Context, arg1 address.Address) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecoverFailedMsg", arg0, arg1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RecoverFailedMsg indicates an expected call of RecoverFailedMsg.
func (mr *MockIMessagerMockRecorder) RecoverFailedMsg(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecoverFailedMsg", reflect.TypeOf((*MockIMessager)(nil).RecoverFailedMsg), arg0, arg1)
}

// RefreshSharedParams mocks base method.
func (m *MockIMessager) RefreshSharedParams(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshSharedParams", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RefreshSharedParams indicates an expected call of RefreshSharedParams.
func (mr *MockIMessagerMockRecorder) RefreshSharedParams(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshSharedParams", reflect.TypeOf((*MockIMessager)(nil).RefreshSharedParams), arg0)
}

// ReplaceMessage mocks base method.
func (m *MockIMessager) ReplaceMessage(arg0 context.Context, arg1 *messager.ReplacMessageParams) (cid.Cid, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReplaceMessage", arg0, arg1)
	ret0, _ := ret[0].(cid.Cid)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReplaceMessage indicates an expected call of ReplaceMessage.
func (mr *MockIMessagerMockRecorder) ReplaceMessage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReplaceMessage", reflect.TypeOf((*MockIMessager)(nil).ReplaceMessage), arg0, arg1)
}

// RepublishMessage mocks base method.
func (m *MockIMessager) RepublishMessage(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RepublishMessage", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RepublishMessage indicates an expected call of RepublishMessage.
func (mr *MockIMessagerMockRecorder) RepublishMessage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RepublishMessage", reflect.TypeOf((*MockIMessager)(nil).RepublishMessage), arg0, arg1)
}

// SaveNode mocks base method.
func (m *MockIMessager) SaveNode(arg0 context.Context, arg1 *messager.Node) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveNode", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveNode indicates an expected call of SaveNode.
func (mr *MockIMessagerMockRecorder) SaveNode(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveNode", reflect.TypeOf((*MockIMessager)(nil).SaveNode), arg0, arg1)
}

// Send mocks base method.
func (m *MockIMessager) Send(arg0 context.Context, arg1 messager.QuickSendParams) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Send indicates an expected call of Send.
func (mr *MockIMessagerMockRecorder) Send(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockIMessager)(nil).Send), arg0, arg1)
}

// SetFeeParams mocks base method.
func (m *MockIMessager) SetFeeParams(arg0 context.Context, arg1 address.Address, arg2, arg3 float64, arg4, arg5 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetFeeParams", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetFeeParams indicates an expected call of SetFeeParams.
func (mr *MockIMessagerMockRecorder) SetFeeParams(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFeeParams", reflect.TypeOf((*MockIMessager)(nil).SetFeeParams), arg0, arg1, arg2, arg3, arg4, arg5)
}

// SetLogLevel mocks base method.
func (m *MockIMessager) SetLogLevel(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetLogLevel", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetLogLevel indicates an expected call of SetLogLevel.
func (mr *MockIMessagerMockRecorder) SetLogLevel(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLogLevel", reflect.TypeOf((*MockIMessager)(nil).SetLogLevel), arg0, arg1)
}

// SetSelectMsgNum mocks base method.
func (m *MockIMessager) SetSelectMsgNum(arg0 context.Context, arg1 address.Address, arg2 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetSelectMsgNum", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetSelectMsgNum indicates an expected call of SetSelectMsgNum.
func (mr *MockIMessagerMockRecorder) SetSelectMsgNum(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSelectMsgNum", reflect.TypeOf((*MockIMessager)(nil).SetSelectMsgNum), arg0, arg1, arg2)
}

// SetSharedParams mocks base method.
func (m *MockIMessager) SetSharedParams(arg0 context.Context, arg1 *messager.SharedSpec) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetSharedParams", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetSharedParams indicates an expected call of SetSharedParams.
func (mr *MockIMessagerMockRecorder) SetSharedParams(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSharedParams", reflect.TypeOf((*MockIMessager)(nil).SetSharedParams), arg0, arg1)
}

// UpdateAllFilledMessage mocks base method.
func (m *MockIMessager) UpdateAllFilledMessage(arg0 context.Context) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllFilledMessage", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAllFilledMessage indicates an expected call of UpdateAllFilledMessage.
func (mr *MockIMessagerMockRecorder) UpdateAllFilledMessage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllFilledMessage", reflect.TypeOf((*MockIMessager)(nil).UpdateAllFilledMessage), arg0)
}

// UpdateFilledMessageByID mocks base method.
func (m *MockIMessager) UpdateFilledMessageByID(arg0 context.Context, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFilledMessageByID", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateFilledMessageByID indicates an expected call of UpdateFilledMessageByID.
func (mr *MockIMessagerMockRecorder) UpdateFilledMessageByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFilledMessageByID", reflect.TypeOf((*MockIMessager)(nil).UpdateFilledMessageByID), arg0, arg1)
}

// UpdateMessageStateByID mocks base method.
func (m *MockIMessager) UpdateMessageStateByID(arg0 context.Context, arg1 string, arg2 messager.MessageState) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMessageStateByID", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMessageStateByID indicates an expected call of UpdateMessageStateByID.
func (mr *MockIMessagerMockRecorder) UpdateMessageStateByID(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMessageStateByID", reflect.TypeOf((*MockIMessager)(nil).UpdateMessageStateByID), arg0, arg1, arg2)
}

// UpdateNonce mocks base method.
func (m *MockIMessager) UpdateNonce(arg0 context.Context, arg1 address.Address, arg2 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateNonce", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateNonce indicates an expected call of UpdateNonce.
func (mr *MockIMessagerMockRecorder) UpdateNonce(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateNonce", reflect.TypeOf((*MockIMessager)(nil).UpdateNonce), arg0, arg1, arg2)
}

// Version mocks base method.
func (m *MockIMessager) Version(arg0 context.Context) (types.Version, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Version", arg0)
	ret0, _ := ret[0].(types.Version)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Version indicates an expected call of Version.
func (mr *MockIMessagerMockRecorder) Version(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockIMessager)(nil).Version), arg0)
}

// WaitMessage mocks base method.
func (m *MockIMessager) WaitMessage(arg0 context.Context, arg1 string, arg2 uint64) (*messager.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitMessage", arg0, arg1, arg2)
	ret0, _ := ret[0].(*messager.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WaitMessage indicates an expected call of WaitMessage.
func (mr *MockIMessagerMockRecorder) WaitMessage(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitMessage", reflect.TypeOf((*MockIMessager)(nil).WaitMessage), arg0, arg1, arg2)
}

// WalletHas mocks base method.
func (m *MockIMessager) WalletHas(arg0 context.Context, arg1 address.Address) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WalletHas", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WalletHas indicates an expected call of WalletHas.
func (mr *MockIMessagerMockRecorder) WalletHas(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WalletHas", reflect.TypeOf((*MockIMessager)(nil).WalletHas), arg0, arg1)
}
