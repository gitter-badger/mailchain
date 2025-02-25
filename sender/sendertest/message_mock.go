// Code generated by MockGen. DO NOT EDIT.
// Source: message.go

// Package sendertest is a generated GoMock package.
package sendertest

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	signer "github.com/mailchain/mailchain/internal/mailbox/signer"
	sender "github.com/mailchain/mailchain/sender"
	reflect "reflect"
)

// MockMessage is a mock of Message interface
type MockMessage struct {
	ctrl     *gomock.Controller
	recorder *MockMessageMockRecorder
}

// MockMessageMockRecorder is the mock recorder for MockMessage
type MockMessageMockRecorder struct {
	mock *MockMessage
}

// NewMockMessage creates a new mock instance
func NewMockMessage(ctrl *gomock.Controller) *MockMessage {
	mock := &MockMessage{ctrl: ctrl}
	mock.recorder = &MockMessageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMessage) EXPECT() *MockMessageMockRecorder {
	return m.recorder
}

// Send mocks base method
func (m *MockMessage) Send(ctx context.Context, network string, to, from, data []byte, signer signer.Signer, opts sender.MessageOpts) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", ctx, network, to, from, data, signer, opts)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockMessageMockRecorder) Send(ctx, network, to, from, data, signer, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockMessage)(nil).Send), ctx, network, to, from, data, signer, opts)
}

// MockMessageOpts is a mock of MessageOpts interface
type MockMessageOpts struct {
	ctrl     *gomock.Controller
	recorder *MockMessageOptsMockRecorder
}

// MockMessageOptsMockRecorder is the mock recorder for MockMessageOpts
type MockMessageOptsMockRecorder struct {
	mock *MockMessageOpts
}

// NewMockMessageOpts creates a new mock instance
func NewMockMessageOpts(ctrl *gomock.Controller) *MockMessageOpts {
	mock := &MockMessageOpts{ctrl: ctrl}
	mock.recorder = &MockMessageOptsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMessageOpts) EXPECT() *MockMessageOptsMockRecorder {
	return m.recorder
}
