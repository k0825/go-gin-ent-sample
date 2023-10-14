// Code generated by MockGen. DO NOT EDIT.
// Source: usecase/interfaces/todo_find.go

// Package mock_interfaces is a generated GoMock package.
package mock_interfaces

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	interfaces "github.com/k0825/go-gin-ent-sample/usecase/interfaces"
)

// MockTodoFindUseCaseInterface is a mock of TodoFindUseCaseInterface interface.
type MockTodoFindUseCaseInterface struct {
	ctrl     *gomock.Controller
	recorder *MockTodoFindUseCaseInterfaceMockRecorder
}

// MockTodoFindUseCaseInterfaceMockRecorder is the mock recorder for MockTodoFindUseCaseInterface.
type MockTodoFindUseCaseInterfaceMockRecorder struct {
	mock *MockTodoFindUseCaseInterface
}

// NewMockTodoFindUseCaseInterface creates a new mock instance.
func NewMockTodoFindUseCaseInterface(ctrl *gomock.Controller) *MockTodoFindUseCaseInterface {
	mock := &MockTodoFindUseCaseInterface{ctrl: ctrl}
	mock.recorder = &MockTodoFindUseCaseInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTodoFindUseCaseInterface) EXPECT() *MockTodoFindUseCaseInterfaceMockRecorder {
	return m.recorder
}

// Handle mocks base method.
func (m *MockTodoFindUseCaseInterface) Handle(arg0 context.Context, arg1 interfaces.TodoFindRequest) (*interfaces.TodoFindResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Handle", arg0, arg1)
	ret0, _ := ret[0].(*interfaces.TodoFindResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Handle indicates an expected call of Handle.
func (mr *MockTodoFindUseCaseInterfaceMockRecorder) Handle(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockTodoFindUseCaseInterface)(nil).Handle), arg0, arg1)
}