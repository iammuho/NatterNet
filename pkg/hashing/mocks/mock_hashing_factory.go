// Code generated by MockGen. DO NOT EDIT.
// Source: hashing_factory.go

// Package mockhashingfactory is a generated GoMock package.
package mockhashingfactory

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockHashingFactory is a mock of HashingFactory interface.
type MockHashingFactory struct {
	ctrl     *gomock.Controller
	recorder *MockHashingFactoryMockRecorder
}

// MockHashingFactoryMockRecorder is the mock recorder for MockHashingFactory.
type MockHashingFactoryMockRecorder struct {
	mock *MockHashingFactory
}

// NewMockHashingFactory creates a new mock instance.
func NewMockHashingFactory(ctrl *gomock.Controller) *MockHashingFactory {
	mock := &MockHashingFactory{ctrl: ctrl}
	mock.recorder = &MockHashingFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHashingFactory) EXPECT() *MockHashingFactoryMockRecorder {
	return m.recorder
}

// ComparePassword mocks base method.
func (m *MockHashingFactory) ComparePassword(password, hash string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComparePassword", password, hash)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ComparePassword indicates an expected call of ComparePassword.
func (mr *MockHashingFactoryMockRecorder) ComparePassword(password, hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComparePassword", reflect.TypeOf((*MockHashingFactory)(nil).ComparePassword), password, hash)
}

// HashPassword mocks base method.
func (m *MockHashingFactory) HashPassword(password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HashPassword", password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HashPassword indicates an expected call of HashPassword.
func (mr *MockHashingFactoryMockRecorder) HashPassword(password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HashPassword", reflect.TypeOf((*MockHashingFactory)(nil).HashPassword), password)
}