// Code generated by MockGen. DO NOT EDIT.
// Source: user_repository.go

// Package mockuserrepository is a generated GoMock package.
package mockuserrepository

import (
	reflect "reflect"

	values "github.com/iammuho/natternet/internal/user/domain/values"
	errorhandler "github.com/iammuho/natternet/pkg/errorhandler"
	gomock "go.uber.org/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUserRepository) Create(user *values.UserDBValue) *errorhandler.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", user)
	ret0, _ := ret[0].(*errorhandler.Response)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockUserRepositoryMockRecorder) Create(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserRepository)(nil).Create), user)
}

// FindOneByEmail mocks base method.
func (m *MockUserRepository) FindOneByEmail(email string) (*values.UserDBValue, *errorhandler.Response) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOneByEmail", email)
	ret0, _ := ret[0].(*values.UserDBValue)
	ret1, _ := ret[1].(*errorhandler.Response)
	return ret0, ret1
}

// FindOneByEmail indicates an expected call of FindOneByEmail.
func (mr *MockUserRepositoryMockRecorder) FindOneByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneByEmail", reflect.TypeOf((*MockUserRepository)(nil).FindOneByEmail), email)
}

// FindOneByID mocks base method.
func (m *MockUserRepository) FindOneByID(id string) (*values.UserDBValue, *errorhandler.Response) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOneByID", id)
	ret0, _ := ret[0].(*values.UserDBValue)
	ret1, _ := ret[1].(*errorhandler.Response)
	return ret0, ret1
}

// FindOneByID indicates an expected call of FindOneByID.
func (mr *MockUserRepositoryMockRecorder) FindOneByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneByID", reflect.TypeOf((*MockUserRepository)(nil).FindOneByID), id)
}

// FindOneByLogin mocks base method.
func (m *MockUserRepository) FindOneByLogin(login string) (*values.UserDBValue, *errorhandler.Response) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOneByLogin", login)
	ret0, _ := ret[0].(*values.UserDBValue)
	ret1, _ := ret[1].(*errorhandler.Response)
	return ret0, ret1
}

// FindOneByLogin indicates an expected call of FindOneByLogin.
func (mr *MockUserRepositoryMockRecorder) FindOneByLogin(login interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneByLogin", reflect.TypeOf((*MockUserRepository)(nil).FindOneByLogin), login)
}

// FindOneByUsername mocks base method.
func (m *MockUserRepository) FindOneByUsername(username string) (*values.UserDBValue, *errorhandler.Response) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOneByUsername", username)
	ret0, _ := ret[0].(*values.UserDBValue)
	ret1, _ := ret[1].(*errorhandler.Response)
	return ret0, ret1
}

// FindOneByUsername indicates an expected call of FindOneByUsername.
func (mr *MockUserRepositoryMockRecorder) FindOneByUsername(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneByUsername", reflect.TypeOf((*MockUserRepository)(nil).FindOneByUsername), username)
}
