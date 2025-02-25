// Code generated by MockGen. DO NOT EDIT.
// Source: producers.go

// Package settingstest is a generated GoMock package.
package settingstest

import (
	gomock "github.com/golang/mock/gomock"
	mailbox "github.com/mailchain/mailchain/internal/mailbox"
	sender "github.com/mailchain/mailchain/sender"
	reflect "reflect"
)

// MockSupporter is a mock of Supporter interface
type MockSupporter struct {
	ctrl     *gomock.Controller
	recorder *MockSupporterMockRecorder
}

// MockSupporterMockRecorder is the mock recorder for MockSupporter
type MockSupporterMockRecorder struct {
	mock *MockSupporter
}

// NewMockSupporter creates a new mock instance
func NewMockSupporter(ctrl *gomock.Controller) *MockSupporter {
	mock := &MockSupporter{ctrl: ctrl}
	mock.recorder = &MockSupporterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSupporter) EXPECT() *MockSupporterMockRecorder {
	return m.recorder
}

// Supports mocks base method
func (m *MockSupporter) Supports() map[string]bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Supports")
	ret0, _ := ret[0].(map[string]bool)
	return ret0
}

// Supports indicates an expected call of Supports
func (mr *MockSupporterMockRecorder) Supports() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Supports", reflect.TypeOf((*MockSupporter)(nil).Supports))
}

// MockSenderClient is a mock of SenderClient interface
type MockSenderClient struct {
	ctrl     *gomock.Controller
	recorder *MockSenderClientMockRecorder
}

// MockSenderClientMockRecorder is the mock recorder for MockSenderClient
type MockSenderClientMockRecorder struct {
	mock *MockSenderClient
}

// NewMockSenderClient creates a new mock instance
func NewMockSenderClient(ctrl *gomock.Controller) *MockSenderClient {
	mock := &MockSenderClient{ctrl: ctrl}
	mock.recorder = &MockSenderClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSenderClient) EXPECT() *MockSenderClientMockRecorder {
	return m.recorder
}

// Produce mocks base method
func (m *MockSenderClient) Produce() (sender.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Produce")
	ret0, _ := ret[0].(sender.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Produce indicates an expected call of Produce
func (mr *MockSenderClientMockRecorder) Produce() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Produce", reflect.TypeOf((*MockSenderClient)(nil).Produce))
}

// Supports mocks base method
func (m *MockSenderClient) Supports() map[string]bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Supports")
	ret0, _ := ret[0].(map[string]bool)
	return ret0
}

// Supports indicates an expected call of Supports
func (mr *MockSenderClientMockRecorder) Supports() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Supports", reflect.TypeOf((*MockSenderClient)(nil).Supports))
}

// MockReceiverClient is a mock of ReceiverClient interface
type MockReceiverClient struct {
	ctrl     *gomock.Controller
	recorder *MockReceiverClientMockRecorder
}

// MockReceiverClientMockRecorder is the mock recorder for MockReceiverClient
type MockReceiverClientMockRecorder struct {
	mock *MockReceiverClient
}

// NewMockReceiverClient creates a new mock instance
func NewMockReceiverClient(ctrl *gomock.Controller) *MockReceiverClient {
	mock := &MockReceiverClient{ctrl: ctrl}
	mock.recorder = &MockReceiverClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockReceiverClient) EXPECT() *MockReceiverClientMockRecorder {
	return m.recorder
}

// Produce mocks base method
func (m *MockReceiverClient) Produce() (mailbox.Receiver, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Produce")
	ret0, _ := ret[0].(mailbox.Receiver)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Produce indicates an expected call of Produce
func (mr *MockReceiverClientMockRecorder) Produce() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Produce", reflect.TypeOf((*MockReceiverClient)(nil).Produce))
}

// Supports mocks base method
func (m *MockReceiverClient) Supports() map[string]bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Supports")
	ret0, _ := ret[0].(map[string]bool)
	return ret0
}

// Supports indicates an expected call of Supports
func (mr *MockReceiverClientMockRecorder) Supports() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Supports", reflect.TypeOf((*MockReceiverClient)(nil).Supports))
}

// MockPublicKeyFinderClient is a mock of PublicKeyFinderClient interface
type MockPublicKeyFinderClient struct {
	ctrl     *gomock.Controller
	recorder *MockPublicKeyFinderClientMockRecorder
}

// MockPublicKeyFinderClientMockRecorder is the mock recorder for MockPublicKeyFinderClient
type MockPublicKeyFinderClientMockRecorder struct {
	mock *MockPublicKeyFinderClient
}

// NewMockPublicKeyFinderClient creates a new mock instance
func NewMockPublicKeyFinderClient(ctrl *gomock.Controller) *MockPublicKeyFinderClient {
	mock := &MockPublicKeyFinderClient{ctrl: ctrl}
	mock.recorder = &MockPublicKeyFinderClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPublicKeyFinderClient) EXPECT() *MockPublicKeyFinderClientMockRecorder {
	return m.recorder
}

// Produce mocks base method
func (m *MockPublicKeyFinderClient) Produce() (mailbox.PubKeyFinder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Produce")
	ret0, _ := ret[0].(mailbox.PubKeyFinder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Produce indicates an expected call of Produce
func (mr *MockPublicKeyFinderClientMockRecorder) Produce() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Produce", reflect.TypeOf((*MockPublicKeyFinderClient)(nil).Produce))
}

// Supports mocks base method
func (m *MockPublicKeyFinderClient) Supports() map[string]bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Supports")
	ret0, _ := ret[0].(map[string]bool)
	return ret0
}

// Supports indicates an expected call of Supports
func (mr *MockPublicKeyFinderClientMockRecorder) Supports() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Supports", reflect.TypeOf((*MockPublicKeyFinderClient)(nil).Supports))
}

// MockSentClient is a mock of SentClient interface
type MockSentClient struct {
	ctrl     *gomock.Controller
	recorder *MockSentClientMockRecorder
}

// MockSentClientMockRecorder is the mock recorder for MockSentClient
type MockSentClientMockRecorder struct {
	mock *MockSentClient
}

// NewMockSentClient creates a new mock instance
func NewMockSentClient(ctrl *gomock.Controller) *MockSentClient {
	mock := &MockSentClient{ctrl: ctrl}
	mock.recorder = &MockSentClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSentClient) EXPECT() *MockSentClientMockRecorder {
	return m.recorder
}

// Produce mocks base method
func (m *MockSentClient) Produce(client string) (sender.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Produce", client)
	ret0, _ := ret[0].(sender.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Produce indicates an expected call of Produce
func (mr *MockSentClientMockRecorder) Produce(client interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Produce", reflect.TypeOf((*MockSentClient)(nil).Produce), client)
}
