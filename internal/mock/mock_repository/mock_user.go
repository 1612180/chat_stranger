// Code generated by MockGen. DO NOT EDIT.
// Source: user.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	model "github.com/1612180/chat_stranger/internal/model"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockUserRepository is a mock of UserRepository interface
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// Find mocks base method
func (m *MockUserRepository) Find(id int) (*model.User, *model.Credential, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(*model.Credential)
	ret2, _ := ret[2].(bool)
	return ret0, ret1, ret2
}

// Find indicates an expected call of Find
func (mr *MockUserRepositoryMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockUserRepository)(nil).Find), id)
}

// FindByRegisterName mocks base method
func (m *MockUserRepository) FindByRegisterName(n string) (*model.User, *model.Credential, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByRegisterName", n)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(*model.Credential)
	ret2, _ := ret[2].(bool)
	return ret0, ret1, ret2
}

// FindByRegisterName indicates an expected call of FindByRegisterName
func (mr *MockUserRepositoryMockRecorder) FindByRegisterName(n interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByRegisterName", reflect.TypeOf((*MockUserRepository)(nil).FindByRegisterName), n)
}

// Create mocks base method
func (m *MockUserRepository) Create(user *model.User, credential *model.Credential) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", user, credential)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockUserRepositoryMockRecorder) Create(user, credential interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserRepository)(nil).Create), user, credential)
}