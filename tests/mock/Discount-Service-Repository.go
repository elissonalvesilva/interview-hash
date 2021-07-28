package mock

import (
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockDiscountServiceRepository is a mock of DiscountServiceRepository interface.
type MockDiscountServiceRepository struct {
	ctrl     *gomock.Controller
	recorder *MockDiscountServiceRepositoryMockRecorder
}

// MockDiscountServiceRepositoryMockRecorder is the mock recorder for MockDiscountServiceRepository.
type MockDiscountServiceRepositoryMockRecorder struct {
	mock *MockDiscountServiceRepository
}

// NewMockDiscountServiceRepository creates a new mock instance.
func NewMockDiscountServiceRepository(ctrl *gomock.Controller) *MockDiscountServiceRepository {
	mock := &MockDiscountServiceRepository{ctrl: ctrl}
	mock.recorder = &MockDiscountServiceRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDiscountServiceRepository) EXPECT() *MockDiscountServiceRepositoryMockRecorder {
	return m.recorder
}

// GetProductDiscount mocks base method.
func (m *MockDiscountServiceRepository) GetProductDiscount(productID int) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductDiscount", productID)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductDiscount indicates an expected call of GetProductDiscount.
func (mr *MockDiscountServiceRepositoryMockRecorder) GetProductDiscount(productID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductDiscount", reflect.TypeOf((*MockDiscountServiceRepository)(nil).GetProductDiscount), productID)
}

