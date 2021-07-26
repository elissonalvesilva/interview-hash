package mock


import (
	"reflect"

	"github.com/elissonalvesilva/interview-hash/my-ecommerce/domain/entity"
	"github.com/elissonalvesilva/interview-hash/my-ecommerce/domain/protocols"
	"github.com/golang/mock/gomock"
)

// MockProductCheckoutRepository is a mock of ProductCheckoutRepository interface.
type MockProductCheckoutRepository struct {
	ctrl     *gomock.Controller
	recorder *MockProductCheckoutRepositoryMockRecorder
}

// MockProductCheckoutRepositoryMockRecorder is the mock recorder for MockProductCheckoutRepository.
type MockProductCheckoutRepositoryMockRecorder struct {
	mock *MockProductCheckoutRepository
}

// NewMockProductCheckoutRepository creates a new mock instance.
func NewMockProductCheckoutRepository(ctrl *gomock.Controller) *MockProductCheckoutRepository {
	mock := &MockProductCheckoutRepository{ctrl: ctrl}
	mock.recorder = &MockProductCheckoutRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductCheckoutRepository) EXPECT() *MockProductCheckoutRepositoryMockRecorder {
	return m.recorder
}

// GetProducts mocks base method.
func (m *MockProductCheckoutRepository) GetProducts(arg0 []protocols.ProductCheckout) []entity.Product {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProducts", arg0)
	ret0, _ := ret[0].([]entity.Product)
	return ret0
}

// GetProducts indicates an expected call of GetProducts.
func (mr *MockProductCheckoutRepositoryMockRecorder) GetProducts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProducts", reflect.TypeOf((*MockProductCheckoutRepository)(nil).GetProducts), arg0)
}

