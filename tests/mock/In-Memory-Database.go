package mock

import (
	reflect "reflect"

	entity "github.com/elissonalvesilva/interview-hash/my-ecommerce/domain/entity"
	protocols "github.com/elissonalvesilva/interview-hash/my-ecommerce/infrastructure/db/in-memory/protocols"
	gomock "github.com/golang/mock/gomock"
)

// MockInMemory is a mock of InMemory interface.
type MockInMemory struct {
	ctrl     *gomock.Controller
	recorder *MockInMemoryMockRecorder
}

// MockInMemoryMockRecorder is the mock recorder for MockInMemory.
type MockInMemoryMockRecorder struct {
	mock *MockInMemory
}

// NewMockInMemory creates a new mock instance.
func NewMockInMemory(ctrl *gomock.Controller) *MockInMemory {
	mock := &MockInMemory{ctrl: ctrl}
	mock.recorder = &MockInMemoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInMemory) EXPECT() *MockInMemoryMockRecorder {
	return m.recorder
}

// GetByID mocks base method.
func (m *MockInMemory) GetByID(id int) (entity.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(entity.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockInMemoryMockRecorder) GetByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockInMemory)(nil).GetByID), id)
}

// GetOne mocks base method.
func (m *MockInMemory) GetOne(filter protocols.Filter) (entity.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOne", filter)
	ret0, _ := ret[0].(entity.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOne indicates an expected call of GetOne.
func (mr *MockInMemoryMockRecorder) GetOne(filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOne", reflect.TypeOf((*MockInMemory)(nil).GetOne), filter)
}
