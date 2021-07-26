package use_cases

import (
	"github.com/elissonalvesilva/interview-hash/my-ecommerce/domain/protocols"
	"github.com/elissonalvesilva/interview-hash/my-ecommerce/tests/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

var productsCheckout = []protocols.ProductCheckout{
	protocols.ProductCheckout{
		ID: 1,
		Quantity: 1,
	},
	protocols.ProductCheckout{
		ID: 2,
		Quantity: 2,
	},
}

func TestProductCheckoutUseCase_CheckoutProducts(t *testing.T) {
	t.Run("Should return a empty response of products", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockProductCheckoutRepository := mock.NewMockProductCheckoutRepository(ctrl)
		mockProductCheckoutRepository.EXPECT().GetProducts(gomock.Any()).Return([]protocols.ProductToApplyDiscount{})

		sut := NewProductsCheckout(mockProductCheckoutRepository)

		response := sut.CheckoutProducts(productsCheckout)
		assert.Equal(t, protocols.CheckoutResponse{}, response)
	})

	t.Run("Should checkout response with correct values", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockProductCheckoutRepository := mock.NewMockProductCheckoutRepository(ctrl)
		mockProductCheckoutRepository.EXPECT().GetProducts(gomock.Any()).Return(mock.ProductsToApplyDiscountResponse)

		sut := NewProductsCheckout(mockProductCheckoutRepository)

		response := sut.CheckoutProducts(productsCheckout)
		assert.Equal(t, mock.CheckoutResponse, response)
	})
}
