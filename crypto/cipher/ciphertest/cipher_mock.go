// Code generated by MockGen. DO NOT EDIT.
// Source: cipher.go

// Package ciphertest is a generated GoMock package.
package ciphertest

import (
	gomock "github.com/golang/mock/gomock"
	crypto "github.com/mailchain/mailchain/crypto"
	cipher "github.com/mailchain/mailchain/crypto/cipher"
	io "io"
	reflect "reflect"
)

// MockDecrypter is a mock of Decrypter interface
type MockDecrypter struct {
	ctrl     *gomock.Controller
	recorder *MockDecrypterMockRecorder
}

// MockDecrypterMockRecorder is the mock recorder for MockDecrypter
type MockDecrypterMockRecorder struct {
	mock *MockDecrypter
}

// NewMockDecrypter creates a new mock instance
func NewMockDecrypter(ctrl *gomock.Controller) *MockDecrypter {
	mock := &MockDecrypter{ctrl: ctrl}
	mock.recorder = &MockDecrypterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDecrypter) EXPECT() *MockDecrypterMockRecorder {
	return m.recorder
}

// Decrypt mocks base method
func (m *MockDecrypter) Decrypt(arg0 cipher.EncryptedContent) (cipher.PlainContent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decrypt", arg0)
	ret0, _ := ret[0].(cipher.PlainContent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Decrypt indicates an expected call of Decrypt
func (mr *MockDecrypterMockRecorder) Decrypt(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decrypt", reflect.TypeOf((*MockDecrypter)(nil).Decrypt), arg0)
}

// MockEncrypter is a mock of Encrypter interface
type MockEncrypter struct {
	ctrl     *gomock.Controller
	recorder *MockEncrypterMockRecorder
}

// MockEncrypterMockRecorder is the mock recorder for MockEncrypter
type MockEncrypterMockRecorder struct {
	mock *MockEncrypter
}

// NewMockEncrypter creates a new mock instance
func NewMockEncrypter(ctrl *gomock.Controller) *MockEncrypter {
	mock := &MockEncrypter{ctrl: ctrl}
	mock.recorder = &MockEncrypterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEncrypter) EXPECT() *MockEncrypterMockRecorder {
	return m.recorder
}

// Encrypt mocks base method
func (m *MockEncrypter) Encrypt(rand io.Reader, pub crypto.PublicKey, plain cipher.PlainContent) (cipher.EncryptedContent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Encrypt", rand, pub, plain)
	ret0, _ := ret[0].(cipher.EncryptedContent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Encrypt indicates an expected call of Encrypt
func (mr *MockEncrypterMockRecorder) Encrypt(rand, pub, plain interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Encrypt", reflect.TypeOf((*MockEncrypter)(nil).Encrypt), rand, pub, plain)
}
