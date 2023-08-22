// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package nats_mock is a generated GoMock package.
package nats_mock

import (
	reflect "reflect"

	nats "github.com/nats-io/nats.go"
	gomock "go.uber.org/mock/gomock"
)

// MockNatsContext is a mock of NatsContext interface.
type MockNatsContext struct {
	ctrl     *gomock.Controller
	recorder *MockNatsContextMockRecorder
}

// MockNatsContextMockRecorder is the mock recorder for MockNatsContext.
type MockNatsContextMockRecorder struct {
	mock *MockNatsContext
}

// NewMockNatsContext creates a new mock instance.
func NewMockNatsContext(ctrl *gomock.Controller) *MockNatsContext {
	mock := &MockNatsContext{ctrl: ctrl}
	mock.recorder = &MockNatsContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNatsContext) EXPECT() *MockNatsContextMockRecorder {
	return m.recorder
}

// CreateStream mocks base method.
func (m *MockNatsContext) CreateStream(streamName, subject string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStream", streamName, subject)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateStream indicates an expected call of CreateStream.
func (mr *MockNatsContextMockRecorder) CreateStream(streamName, subject interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStream", reflect.TypeOf((*MockNatsContext)(nil).CreateStream), streamName, subject)
}

// GetConn mocks base method.
func (m *MockNatsContext) GetConn() *nats.Conn {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConn")
	ret0, _ := ret[0].(*nats.Conn)
	return ret0
}

// GetConn indicates an expected call of GetConn.
func (mr *MockNatsContextMockRecorder) GetConn() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConn", reflect.TypeOf((*MockNatsContext)(nil).GetConn))
}

// GetJetStreamContext mocks base method.
func (m *MockNatsContext) GetJetStreamContext() nats.JetStreamContext {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJetStreamContext")
	ret0, _ := ret[0].(nats.JetStreamContext)
	return ret0
}

// GetJetStreamContext indicates an expected call of GetJetStreamContext.
func (mr *MockNatsContextMockRecorder) GetJetStreamContext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJetStreamContext", reflect.TypeOf((*MockNatsContext)(nil).GetJetStreamContext))
}

// Subscribe mocks base method.
func (m *MockNatsContext) Subscribe(subject string, handler func(*nats.Msg) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", subject, handler)
	ret0, _ := ret[0].(error)
	return ret0
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockNatsContextMockRecorder) Subscribe(subject, handler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockNatsContext)(nil).Subscribe), subject, handler)
}
