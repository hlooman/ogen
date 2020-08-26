// Code generated by MockGen. DO NOT EDIT.
// Source: internal/hostnode/hostnode.go

// Package hostnode is a generated GoMock package.
package hostnode

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	host "github.com/libp2p/go-libp2p-core/host"
	network "github.com/libp2p/go-libp2p-core/network"
	peer "github.com/libp2p/go-libp2p-core/peer"
	protocol "github.com/libp2p/go-libp2p-core/protocol"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	multiaddr "github.com/multiformats/go-multiaddr"
	reflect "reflect"
)

// MockHostNode is a mock of HostNode interface
type MockHostNode struct {
	ctrl     *gomock.Controller
	recorder *MockHostNodeMockRecorder
}

// MockHostNodeMockRecorder is the mock recorder for MockHostNode
type MockHostNodeMockRecorder struct {
	mock *MockHostNode
}

// NewMockHostNode creates a new mock instance
func NewMockHostNode(ctrl *gomock.Controller) *MockHostNode {
	mock := &MockHostNode{ctrl: ctrl}
	mock.recorder = &MockHostNodeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHostNode) EXPECT() *MockHostNodeMockRecorder {
	return m.recorder
}

// SyncProtocol mocks base method
func (m *MockHostNode) SyncProtocol() SyncProtocol {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncProtocol")
	ret0, _ := ret[0].(SyncProtocol)
	return ret0
}

// SyncProtocol indicates an expected call of SyncProtocol
func (mr *MockHostNodeMockRecorder) SyncProtocol() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncProtocol", reflect.TypeOf((*MockHostNode)(nil).SyncProtocol))
}

// Topic mocks base method
func (m *MockHostNode) Topic(topic string) (*pubsub.Topic, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Topic", topic)
	ret0, _ := ret[0].(*pubsub.Topic)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Topic indicates an expected call of Topic
func (mr *MockHostNodeMockRecorder) Topic(topic interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Topic", reflect.TypeOf((*MockHostNode)(nil).Topic), topic)
}

// Syncing mocks base method
func (m *MockHostNode) Syncing() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Syncing")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Syncing indicates an expected call of Syncing
func (mr *MockHostNodeMockRecorder) Syncing() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Syncing", reflect.TypeOf((*MockHostNode)(nil).Syncing))
}

// GetContext mocks base method
func (m *MockHostNode) GetContext() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContext")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// GetContext indicates an expected call of GetContext
func (mr *MockHostNodeMockRecorder) GetContext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContext", reflect.TypeOf((*MockHostNode)(nil).GetContext))
}

// GetHost mocks base method
func (m *MockHostNode) GetHost() host.Host {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHost")
	ret0, _ := ret[0].(host.Host)
	return ret0
}

// GetHost indicates an expected call of GetHost
func (mr *MockHostNodeMockRecorder) GetHost() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHost", reflect.TypeOf((*MockHostNode)(nil).GetHost))
}

// GetNetMagic mocks base method
func (m *MockHostNode) GetNetMagic() uint32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNetMagic")
	ret0, _ := ret[0].(uint32)
	return ret0
}

// GetNetMagic indicates an expected call of GetNetMagic
func (mr *MockHostNodeMockRecorder) GetNetMagic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNetMagic", reflect.TypeOf((*MockHostNode)(nil).GetNetMagic))
}

// removePeer mocks base method
func (m *MockHostNode) removePeer(p peer.ID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "removePeer", p)
}

// removePeer indicates an expected call of removePeer
func (mr *MockHostNodeMockRecorder) removePeer(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "removePeer", reflect.TypeOf((*MockHostNode)(nil).removePeer), p)
}

// DisconnectPeer mocks base method
func (m *MockHostNode) DisconnectPeer(p peer.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisconnectPeer", p)
	ret0, _ := ret[0].(error)
	return ret0
}

// DisconnectPeer indicates an expected call of DisconnectPeer
func (mr *MockHostNodeMockRecorder) DisconnectPeer(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisconnectPeer", reflect.TypeOf((*MockHostNode)(nil).DisconnectPeer), p)
}

// IsConnected mocks base method
func (m *MockHostNode) IsConnected() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsConnected")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsConnected indicates an expected call of IsConnected
func (mr *MockHostNodeMockRecorder) IsConnected() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsConnected", reflect.TypeOf((*MockHostNode)(nil).IsConnected))
}

// PeersConnected mocks base method
func (m *MockHostNode) PeersConnected() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PeersConnected")
	ret0, _ := ret[0].(int)
	return ret0
}

// PeersConnected indicates an expected call of PeersConnected
func (mr *MockHostNodeMockRecorder) PeersConnected() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PeersConnected", reflect.TypeOf((*MockHostNode)(nil).PeersConnected))
}

// GetPeerList mocks base method
func (m *MockHostNode) GetPeerList() []peer.ID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPeerList")
	ret0, _ := ret[0].([]peer.ID)
	return ret0
}

// GetPeerList indicates an expected call of GetPeerList
func (mr *MockHostNodeMockRecorder) GetPeerList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPeerList", reflect.TypeOf((*MockHostNode)(nil).GetPeerList))
}

// GetPeerInfos mocks base method
func (m *MockHostNode) GetPeerInfos() []peer.AddrInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPeerInfos")
	ret0, _ := ret[0].([]peer.AddrInfo)
	return ret0
}

// GetPeerInfos indicates an expected call of GetPeerInfos
func (mr *MockHostNodeMockRecorder) GetPeerInfos() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPeerInfos", reflect.TypeOf((*MockHostNode)(nil).GetPeerInfos))
}

// ConnectedToPeer mocks base method
func (m *MockHostNode) ConnectedToPeer(id peer.ID) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectedToPeer", id)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ConnectedToPeer indicates an expected call of ConnectedToPeer
func (mr *MockHostNodeMockRecorder) ConnectedToPeer(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectedToPeer", reflect.TypeOf((*MockHostNode)(nil).ConnectedToPeer), id)
}

// Notify mocks base method
func (m *MockHostNode) Notify(notifee network.Notifiee) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Notify", notifee)
}

// Notify indicates an expected call of Notify
func (mr *MockHostNodeMockRecorder) Notify(notifee interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Notify", reflect.TypeOf((*MockHostNode)(nil).Notify), notifee)
}

// setStreamHandler mocks base method
func (m *MockHostNode) setStreamHandler(id protocol.ID, handleStream func(network.Stream)) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "setStreamHandler", id, handleStream)
}

// setStreamHandler indicates an expected call of setStreamHandler
func (mr *MockHostNodeMockRecorder) setStreamHandler(id, handleStream interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "setStreamHandler", reflect.TypeOf((*MockHostNode)(nil).setStreamHandler), id, handleStream)
}

// CountPeers mocks base method
func (m *MockHostNode) CountPeers(id protocol.ID) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountPeers", id)
	ret0, _ := ret[0].(int)
	return ret0
}

// CountPeers indicates an expected call of CountPeers
func (mr *MockHostNodeMockRecorder) CountPeers(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountPeers", reflect.TypeOf((*MockHostNode)(nil).CountPeers), id)
}

// GetPeerDirection mocks base method
func (m *MockHostNode) GetPeerDirection(id peer.ID) network.Direction {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPeerDirection", id)
	ret0, _ := ret[0].(network.Direction)
	return ret0
}

// GetPeerDirection indicates an expected call of GetPeerDirection
func (mr *MockHostNodeMockRecorder) GetPeerDirection(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPeerDirection", reflect.TypeOf((*MockHostNode)(nil).GetPeerDirection), id)
}

// Start mocks base method
func (m *MockHostNode) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockHostNodeMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockHostNode)(nil).Start))
}

// SavePeer mocks base method
func (m *MockHostNode) SavePeer(pma multiaddr.Multiaddr) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SavePeer", pma)
	ret0, _ := ret[0].(error)
	return ret0
}

// SavePeer indicates an expected call of SavePeer
func (mr *MockHostNodeMockRecorder) SavePeer(pma interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SavePeer", reflect.TypeOf((*MockHostNode)(nil).SavePeer), pma)
}

// BanScorePeer mocks base method
func (m *MockHostNode) BanScorePeer(id peer.ID, weight int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BanScorePeer", id, weight)
	ret0, _ := ret[0].(error)
	return ret0
}

// BanScorePeer indicates an expected call of BanScorePeer
func (mr *MockHostNodeMockRecorder) BanScorePeer(id, weight interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BanScorePeer", reflect.TypeOf((*MockHostNode)(nil).BanScorePeer), id, weight)
}

// IsPeerBanned mocks base method
func (m *MockHostNode) IsPeerBanned(id peer.ID) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsPeerBanned", id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsPeerBanned indicates an expected call of IsPeerBanned
func (mr *MockHostNodeMockRecorder) IsPeerBanned(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsPeerBanned", reflect.TypeOf((*MockHostNode)(nil).IsPeerBanned), id)
}