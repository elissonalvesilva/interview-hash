package use_cases

import (
	"errors"
	"github.com/elissonalvesilva/interview-hash/my-ecommerce/domain/protocols"
	"github.com/elissonalvesilva/interview-hash/my-ecommerce/tests/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
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

var blackFridayDate = time.Date(2021, 11, 29, 00, 00, 00, 00, time.UTC)

func TestProductCheckoutUseCase_CheckoutProducts(t *testing.T) {
	t.Run("Should return a empty response of products", func(t *testing.T) {
		t.Parallel()
		ctrl := gomock.NewController(t)

		mockProductCheckoutRepository := mock.NewMockProductCheckoutRepository(ctrl)
		mockDiscountServiceRepository := mock.NewMockDiscountServiceRepository(ctrl)

		mockProductCheckoutRepository.EXPECT().GetProducts(gomock.Any()).Return([]protocols.ProductToApplyDiscount{})

		sut := NewProductsCheckout(mockProductCheckoutRepository, mockDiscountServiceRepository, blackFridayDate)

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

		sut := NewProductsCheckout(mockProductCheckoutRepository, mockDiscountServiceRepository, blackFridayDate)

		response := sut.CheckoutProducts(productsCheckout)
		assert.Equal(t, mock.CheckoutResponseWithoutDiscount, response)
	})

	t.Run("Should return a response with products applied discount and with a gift 'cause the date is equals to black friday", func(t *testing.T) {
		t.Parallel()
		ctrl := gomock.NewController(t)

		mockProductCheckoutRepository := mock.NewMockProductCheckoutRepository(ctrl)
		mockDiscountServiceRepository := mock.NewMockDiscountServiceRepository(ctrl)

		giftProduct := protocols.ProductAppliedDiscount{
			ID: mock.Product3.ID,
			Quantity: 2,
			UnitAmount: 0,
			TotalAmount: 0,
			Discount: 0,
			IsGift: mock.Product3.IsGift,
		}

		gomock.InOrder(
			mockDiscountServiceRepository.EXPECT().GetProductDiscount(gomock.Any()).Return(0.15, nil),
			mockDiscountServiceRepository.EXPECT().GetProductDiscount(gomock.Any()).Return(0.15, nil),
		)

		mockProductCheckoutRepository.EXPECT().GetProducts(gomock.Any()).Return(mock.ProductsToApplyDiscountWithoutGift)
		mockProductCheckoutRepository.EXPECT().GetProductToGift().Return(giftProduct)

		sut := NewProductsCheckout(mockProductCheckoutRepository, mockDiscountServiceRepository, time.Now())

		expectedResponse := mock.CheckoutResponseWithGift
		expectedResponse.Products = append(expectedResponse.Products, giftProduct)

		response := sut.CheckoutProducts(productsCheckout)
		assert.Equal(t, expectedResponse, response)
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

		sut := NewProductsCheckout(mockProductCheckoutRepository, mockDiscountServiceRepository, blackFridayDate)

		response := sut.CheckoutProducts(productsCheckout)
		assert.Equal(t, mock.CheckoutResponse, response)
	})
}

func TestExistsGiftAddedInProducts(t *testing.T) {
	t.Run("Should return false if not exists product gift in list", func(t *testing.T) {
		t.Parallel()
		expectedResponse := false

		existsGiftInList := ExistsGiftAddedInProducts(mock.ProductsAppliedDiscountWithoutGift)
		assert.Equal(t, expectedResponse, existsGiftInList)
	})

	t.Run("Should return true if exists product gift in list", func(t *testing.T) {
		t.Parallel()
		expectedResponse := true

		existsGiftInList := ExistsGiftAddedInProducts(mock.ProductsAppliedDiscount)
		assert.Equal(t, expectedResponse, existsGiftInList)
	})
}
