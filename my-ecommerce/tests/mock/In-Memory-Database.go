package mock

import (
	"reflect"

	"github.com/elissonalvesilva/interview-hash/my-ecommerce/domain/entity"
	"github.com/golang/mock/gomock"
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
