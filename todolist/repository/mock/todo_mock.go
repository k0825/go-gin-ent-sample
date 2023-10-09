// Code generated by MockGen. DO NOT EDIT.
// Source: repository/interfaces/todo.go

// Package mock_interfaces is a generated GoMock package.
package mock_interfaces

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	models "github.com/k0825/go-gin-ent-sample/models"
)

// MockTodoRepositoryInterface is a mock of TodoRepositoryInterface interface.
type MockTodoRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockTodoRepositoryInterfaceMockRecorder
}

// MockTodoRepositoryInterfaceMockRecorder is the mock recorder for MockTodoRepositoryInterface.
type MockTodoRepositoryInterfaceMockRecorder struct {
	mock *MockTodoRepositoryInterface
}

// NewMockTodoRepositoryInterface creates a new mock instance.
func NewMockTodoRepositoryInterface(ctrl *gomock.Controller) *MockTodoRepositoryInterface {
	mock := &MockTodoRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockTodoRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTodoRepositoryInterface) EXPECT() *MockTodoRepositoryInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockTodoRepositoryInterface) Create(arg0 context.Context, arg1 models.TodoTitle, arg2 models.TodoDescription, arg3 models.TodoImage, arg4 []models.TodoTag, arg5, arg6 time.Time) (*models.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(*models.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockTodoRepositoryInterfaceMockRecorder) Create(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTodoRepositoryInterface)(nil).Create), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// FindById mocks base method.
func (m *MockTodoRepositoryInterface) FindById(arg0 context.Context, arg1 models.TodoId) (*models.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", arg0, arg1)
	ret0, _ := ret[0].(*models.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockTodoRepositoryInterfaceMockRecorder) FindById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockTodoRepositoryInterface)(nil).FindById), arg0, arg1)
}
