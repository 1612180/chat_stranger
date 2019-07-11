// Code generated by MockGen. DO NOT EDIT.
// Source: chat.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	model "github.com/1612180/chat_stranger/internal/model"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockChatService is a mock of ChatService interface
type MockChatService struct {
	ctrl     *gomock.Controller
	recorder *MockChatServiceMockRecorder
}

// MockChatServiceMockRecorder is the mock recorder for MockChatService
type MockChatServiceMockRecorder struct {
	mock *MockChatService
}

// NewMockChatService creates a new mock instance
func NewMockChatService(ctrl *gomock.Controller) *MockChatService {
	mock := &MockChatService{ctrl: ctrl}
	mock.recorder = &MockChatServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockChatService) EXPECT() *MockChatServiceMockRecorder {
	return m.recorder
}

// Find mocks base method
func (m *MockChatService) Find(userID int, status string) (*model.Room, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", userID, status)
	ret0, _ := ret[0].(*model.Room)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockChatServiceMockRecorder) Find(userID, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockChatService)(nil).Find), userID, status)
}

// Join mocks base method
func (m *MockChatService) Join(userID, roomID int) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Join", userID, roomID)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Join indicates an expected call of Join
func (mr *MockChatServiceMockRecorder) Join(userID, roomID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Join", reflect.TypeOf((*MockChatService)(nil).Join), userID, roomID)
}

// Leave mocks base method
func (m *MockChatService) Leave(userID int) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Leave", userID)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Leave indicates an expected call of Leave
func (mr *MockChatServiceMockRecorder) Leave(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Leave", reflect.TypeOf((*MockChatService)(nil).Leave), userID)
}

// SendMessage mocks base method
func (m *MockChatService) SendMessage(message *model.Message) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMessage", message)
	ret0, _ := ret[0].(bool)
	return ret0
}

// SendMessage indicates an expected call of SendMessage
func (mr *MockChatServiceMockRecorder) SendMessage(message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessage", reflect.TypeOf((*MockChatService)(nil).SendMessage), message)
}

// ReceiveMessage mocks base method
func (m *MockChatService) ReceiveMessage(userID int, fromTime time.Time) ([]*model.Message, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReceiveMessage", userID, fromTime)
	ret0, _ := ret[0].([]*model.Message)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// ReceiveMessage indicates an expected call of ReceiveMessage
func (mr *MockChatServiceMockRecorder) ReceiveMessage(userID, fromTime interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReceiveMessage", reflect.TypeOf((*MockChatService)(nil).ReceiveMessage), userID, fromTime)
}

// IsUserFree mocks base method
func (m *MockChatService) IsUserFree(userID int) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsUserFree", userID)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsUserFree indicates an expected call of IsUserFree
func (mr *MockChatServiceMockRecorder) IsUserFree(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUserFree", reflect.TypeOf((*MockChatService)(nil).IsUserFree), userID)
}

// CountMember mocks base method
func (m *MockChatService) CountMember(userID int) (int, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountMember", userID)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// CountMember indicates an expected call of CountMember
func (mr *MockChatServiceMockRecorder) CountMember(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountMember", reflect.TypeOf((*MockChatService)(nil).CountMember), userID)
}