package use_cases

import (
	"errors"
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
		t.Parallel()
		ctrl := gomock.NewController(t)

		mockProductCheckoutRepository := mock.NewMockProductCheckoutRepository(ctrl)
		mockDiscountServiceRepository := mock.NewMockDiscountServiceRepository(ctrl)

		mockProductCheckoutRepository.EXPECT().GetProducts(gomock.Any()).Return([]protocols.ProductToApplyDiscount{})

		sut := NewProductsCheckout(mockProductCheckoutRepository, mockDiscountServiceRepository)

		response := sut.CheckoutProducts(productsCheckout)
		assert.Equal(t, protocols.CheckoutResponse{}, response)
	})

	t.Run("Should apply discount equals to 0 if service discount returns error", func(t *testing.T) {
		t.Parallel()
		ctrl := gomock.NewController(t)

		mockProductCheckoutRepository := mock.NewMockProductCheckoutRepository(ctrl)
		mockDiscountServiceRepository := mock.NewMockDiscountServiceRepository(ctrl)

		gomock.InOrder(
			mockDiscountServiceRepository.EXPECT().GetProductDiscount(gomock.Any()).Return(0.0, errors.New("service unavailable")),
			mockDiscountServiceRepository.EXPECT().GetProductDiscount(gomock.Any()).Return(0.25, nil),
		)
		mockProductCheckoutRepository.EXPECT().GetProducts(gomock.Any()).Return(mock.ProductsToApplyDiscountResponse)

		sut := NewProductsCheckout(mockProductCheckoutRepository, mockDiscountServiceRepository)

		response := sut.CheckoutProducts(productsCheckout)
		assert.Equal(t, mock.CheckoutResponseWithoutDiscount, response)
	})

	t.Run("Should return a response with products applied discount and total of amount, discount and applied discount", func(t *testing.T) {
		t.Parallel()
		ctrl := gomock.NewController(t)

		mockProductCheckoutRepository := mock.NewMockProductCheckoutRepository(ctrl)
		mockDiscountServiceRepository := mock.NewMockDiscountServiceRepository(ctrl)

		gomock.InOrder(
			mockDiscountServiceRepository.EXPECT().GetProductDiscount(gomock.Any()).Return(0.15, nil),
			mockDiscountServiceRepository.EXPECT().GetProductDiscount(gomock.Any()).Return(0.25, nil),
		)
		mockProductCheckoutRepository.EXPECT().GetProducts(gomock.Any()).Return(mock.ProductsToApplyDiscountResponse)

		sut := NewProductsCheckout(mockProductCheckoutRepository, mockDiscountServiceRepository)

		response := sut.CheckoutProducts(productsCheckout)
		assert.Equal(t, mock.CheckoutResponse, response)
	})
}
