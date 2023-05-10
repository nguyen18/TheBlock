// Code generated by MockGen. DO NOT EDIT.
// Source: TheBlock/src/api/auth/internal/datastore (interfaces: AuthDatastore)

// Package mockauthdatastore is a generated GoMock package.
package mockauthdatastore

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAuthDatastore is a mock of AuthDatastore interface.
type MockAuthDatastore struct {
	ctrl     *gomock.Controller
	recorder *MockAuthDatastoreMockRecorder
}

// MockAuthDatastoreMockRecorder is the mock recorder for MockAuthDatastore.
type MockAuthDatastoreMockRecorder struct {
	mock *MockAuthDatastore
}

// NewMockAuthDatastore creates a new mock instance.
func NewMockAuthDatastore(ctrl *gomock.Controller) *MockAuthDatastore {
	mock := &MockAuthDatastore{ctrl: ctrl}
	mock.recorder = &MockAuthDatastoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthDatastore) EXPECT() *MockAuthDatastoreMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockAuthDatastore) CreateUser(arg0 context.Context, arg1, arg2 string, arg3 []byte) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockAuthDatastoreMockRecorder) CreateUser(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockAuthDatastore)(nil).CreateUser), arg0, arg1, arg2, arg3)
}

// GetUserIdByUuid mocks base method.
func (m *MockAuthDatastore) GetUserIdByUuid(arg0 context.Context, arg1 string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserIdByUuid", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserIdByUuid indicates an expected call of GetUserIdByUuid.
func (mr *MockAuthDatastoreMockRecorder) GetUserIdByUuid(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserIdByUuid", reflect.TypeOf((*MockAuthDatastore)(nil).GetUserIdByUuid), arg0, arg1)
}

// GetUserPasswordByEmail mocks base method.
func (m *MockAuthDatastore) GetUserPasswordByEmail(arg0 context.Context, arg1 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserPasswordByEmail", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserPasswordByEmail indicates an expected call of GetUserPasswordByEmail.
func (mr *MockAuthDatastoreMockRecorder) GetUserPasswordByEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserPasswordByEmail", reflect.TypeOf((*MockAuthDatastore)(nil).GetUserPasswordByEmail), arg0, arg1)
}

// GetUserUuidByEmail mocks base method.
func (m *MockAuthDatastore) GetUserUuidByEmail(arg0 context.Context, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserUuidByEmail", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserUuidByEmail indicates an expected call of GetUserUuidByEmail.
func (mr *MockAuthDatastoreMockRecorder) GetUserUuidByEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserUuidByEmail", reflect.TypeOf((*MockAuthDatastore)(nil).GetUserUuidByEmail), arg0, arg1)
}

// UserExists mocks base method.
func (m *MockAuthDatastore) UserExists(arg0 context.Context, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserExists", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserExists indicates an expected call of UserExists.
func (mr *MockAuthDatastoreMockRecorder) UserExists(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserExists", reflect.TypeOf((*MockAuthDatastore)(nil).UserExists), arg0, arg1)
}