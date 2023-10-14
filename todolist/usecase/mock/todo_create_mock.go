// Code generated by MockGen. DO NOT EDIT.
// Source: usecase/interfaces/todo_create.go

// Package mock_interfaces is a generated GoMock package.
package mock_interfaces

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	interfaces "github.com/k0825/go-gin-ent-sample/usecase/interfaces"
)

// MockTodoCreateUseCaseInterface is a mock of TodoCreateUseCaseInterface interface.
type MockTodoCreateUseCaseInterface struct {
	ctrl     *gomock.Controller
	recorder *MockTodoCreateUseCaseInterfaceMockRecorder
}

// MockTodoCreateUseCaseInterfaceMockRecorder is the mock recorder for MockTodoCreateUseCaseInterface.
type MockTodoCreateUseCaseInterfaceMockRecorder struct {
	mock *MockTodoCreateUseCaseInterface
}

// NewMockTodoCreateUseCaseInterface creates a new mock instance.
func NewMockTodoCreateUseCaseInterface(ctrl *gomock.Controller) *MockTodoCreateUseCaseInterface {
	mock := &MockTodoCreateUseCaseInterface{ctrl: ctrl}
	mock.recorder = &MockTodoCreateUseCaseInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTodoCreateUseCaseInterface) EXPECT() *MockTodoCreateUseCaseInterfaceMockRecorder {
	return m.recorder
}

// Handle mocks base method.
func (m *MockTodoCreateUseCaseInterface) Handle(arg0 context.Context, arg1 interfaces.TodoCreateRequest) (*interfaces.TodoCreateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Handle", arg0, arg1)
	ret0, _ := ret[0].(*interfaces.TodoCreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Handle indicates an expected call of Handle.
func (mr *MockTodoCreateUseCaseInterfaceMockRecorder) Handle(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockTodoCreateUseCaseInterface)(nil).Handle), arg0, arg1)
}