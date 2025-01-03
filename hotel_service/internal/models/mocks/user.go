// Code generated by MockGen. DO NOT EDIT.
// Source: hotel_service/internal/models/user.go
//
// Generated by this command:
//
//	mockgen -source=hotel_service/internal/models/user.go -destination=hotel_service/internal/models/mocks/user.go
//

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	models "github.com/WhoDoIt/gofinal/hotel_service/internal/models"
	gomock "go.uber.org/mock/gomock"
)

// MockUserModelInterface is a mock of UserModelInterface interface.
type MockUserModelInterface struct {
	ctrl     *gomock.Controller
	recorder *MockUserModelInterfaceMockRecorder
	isgomock struct{}
}

// MockUserModelInterfaceMockRecorder is the mock recorder for MockUserModelInterface.
type MockUserModelInterfaceMockRecorder struct {
	mock *MockUserModelInterface
}

// NewMockUserModelInterface creates a new mock instance.
func NewMockUserModelInterface(ctrl *gomock.Controller) *MockUserModelInterface {
	mock := &MockUserModelInterface{ctrl: ctrl}
	mock.recorder = &MockUserModelInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserModelInterface) EXPECT() *MockUserModelInterfaceMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockUserModelInterface) Get(ctx context.Context, user_id int) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, user_id)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockUserModelInterfaceMockRecorder) Get(ctx, user_id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUserModelInterface)(nil).Get), ctx, user_id)
}
