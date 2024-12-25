// Code generated by MockGen. DO NOT EDIT.
// Source: hotel_service/internal/models/room.go
//
// Generated by this command:
//
//	mockgen -source=hotel_service/internal/models/room.go -destination=hotel_service/internal/models/mocks/room.go
//

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	models "github.com/WhoDoIt/gofinal/hotel_service/internal/models"
	gomock "go.uber.org/mock/gomock"
)

// MockRoomModelInterface is a mock of RoomModelInterface interface.
type MockRoomModelInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRoomModelInterfaceMockRecorder
	isgomock struct{}
}

// MockRoomModelInterfaceMockRecorder is the mock recorder for MockRoomModelInterface.
type MockRoomModelInterfaceMockRecorder struct {
	mock *MockRoomModelInterface
}

// NewMockRoomModelInterface creates a new mock instance.
func NewMockRoomModelInterface(ctrl *gomock.Controller) *MockRoomModelInterface {
	mock := &MockRoomModelInterface{ctrl: ctrl}
	mock.recorder = &MockRoomModelInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoomModelInterface) EXPECT() *MockRoomModelInterfaceMockRecorder {
	return m.recorder
}

// DeleteRoom mocks base method.
func (m *MockRoomModelInterface) DeleteRoom(ctx context.Context, hotel_id, room_id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRoom", ctx, hotel_id, room_id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRoom indicates an expected call of DeleteRoom.
func (mr *MockRoomModelInterfaceMockRecorder) DeleteRoom(ctx, hotel_id, room_id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRoom", reflect.TypeOf((*MockRoomModelInterface)(nil).DeleteRoom), ctx, hotel_id, room_id)
}

// GetAllInHotel mocks base method.
func (m *MockRoomModelInterface) GetAllInHotel(ctx context.Context, hotel_id int) ([]*models.Room, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllInHotel", ctx, hotel_id)
	ret0, _ := ret[0].([]*models.Room)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllInHotel indicates an expected call of GetAllInHotel.
func (mr *MockRoomModelInterfaceMockRecorder) GetAllInHotel(ctx, hotel_id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllInHotel", reflect.TypeOf((*MockRoomModelInterface)(nil).GetAllInHotel), ctx, hotel_id)
}

// GetById mocks base method.
func (m *MockRoomModelInterface) GetById(ctx context.Context, room_id int) (*models.Room, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", ctx, room_id)
	ret0, _ := ret[0].(*models.Room)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockRoomModelInterfaceMockRecorder) GetById(ctx, room_id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockRoomModelInterface)(nil).GetById), ctx, room_id)
}

// Insert mocks base method.
func (m *MockRoomModelInterface) Insert(ctx context.Context, hotel_id int, room_type string, room_price float32) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, hotel_id, room_type, room_price)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert.
func (mr *MockRoomModelInterfaceMockRecorder) Insert(ctx, hotel_id, room_type, room_price any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockRoomModelInterface)(nil).Insert), ctx, hotel_id, room_type, room_price)
}

// UpdateRoom mocks base method.
func (m *MockRoomModelInterface) UpdateRoom(ctx context.Context, hotel_id, room_id int, room_type string, room_price float32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRoom", ctx, hotel_id, room_id, room_type, room_price)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRoom indicates an expected call of UpdateRoom.
func (mr *MockRoomModelInterfaceMockRecorder) UpdateRoom(ctx, hotel_id, room_id, room_type, room_price any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRoom", reflect.TypeOf((*MockRoomModelInterface)(nil).UpdateRoom), ctx, hotel_id, room_id, room_type, room_price)
}