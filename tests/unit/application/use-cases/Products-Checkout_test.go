package use_cases

import (
	"errors"
	"github.com/elissonalvesilva/interview-hash/my-ecommerce/application/use-cases"
	"github.com/elissonalvesilva/interview-hash/my-ecommerce/domain/protocols"
	mock2 "github.com/elissonalvesilva/interview-hash/tests/mock"
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

		mockProductCheckoutRepository := mock2.NewMockProductCheckoutRepository(ctrl)
		mockDiscountServiceRepository := mock2.NewMockDiscountServiceRepository(ctrl)

		mockProductCheckoutRepository.EXPECT().GetProducts(gomock.Any()).Return([]protocols.ProductToApplyDiscount{})

		sut := use_cases.NewProductsCheckout(mockProductCheckoutRepository, mockDiscountServiceRepository, blackFridayDate)

		response := sut.CheckoutProducts(productsCheckout)
		assert.Equal(t, protocols.CheckoutResponse{}, response)
	})

	t.Run("Should apply discount equals to 0 if service discount returns error", func(t *testing.T) {
		t.Parallel()
		ctrl := gomock.NewController(t)

		mockProductCheckoutRepository := mock2.NewMockProductCheckoutRepository(ctrl)
		mockDiscountServiceRepository := mock2.NewMockDiscountServiceRepository(ctrl)

		gomock.InOrder(
			mockDiscountServiceRepository.EXPECT().GetProductDiscount(gomock.Any()).Return(0.0, errors.New("service unavailable")),
			mockDiscountServiceRepository.EXPECT().GetProductDiscount(gomock.Any()).Return(0.25, nil),
		)
		mockProductCheckoutRepository.EXPECT().GetProducts(gomock.Any()).Return(mock2.ProductsToApplyDiscountResponse)

		sut := use_cases.NewProductsCheckout(mockProductCheckoutRepository, mockDiscountServiceRepository, blackFridayDate)

		response := sut.CheckoutProducts(productsCheckout)
		assert.Equal(t, mock2.CheckoutResponseWithoutDiscount, response)
	})

	t.Run("Should return a response with products applied discount and with a gift 'cause the date is equals to black friday", func(t *testing.T) {
		t.Parallel()
		ctrl := gomock.NewController(t)

		mockProductCheckoutRepository := mock2.NewMockProductCheckoutRepository(ctrl)
		mockDiscountServiceRepository := mock2.NewMockDiscountServiceRepository(ctrl)

		giftProduct := protocols.ProductAppliedDiscount{
			ID:          mock2.Product3.ID,
			Quantity:    2,
			UnitAmount:  0,
			TotalAmount: 0,
			Discount:    0,
			IsGift:      mock2.Product3.IsGift,
		}

		gomock.InOrder(
			mockDiscountServiceRepository.EXPECT().GetProductDiscount(gomock.Any()).Return(0.15, nil),
			mockDiscountServiceRepository.EXPECT().GetProductDiscount(gomock.Any()).Return(0.15, nil),
		)

		mockProductCheckoutRepository.EXPECT().GetProducts(gomock.Any()).Return(mock2.ProductsToApplyDiscountWithoutGift)
		mockProductCheckoutRepository.EXPECT().GetProductToGift().Return(giftProduct)

		sut := use_cases.NewProductsCheckout(mockProductCheckoutRepository, mockDiscountServiceRepository, time.Now())

		expectedResponse := mock2.CheckoutResponseWithGift
		expectedResponse.Products = append(expectedResponse.Products, giftProduct)

		response := sut.CheckoutProducts(productsCheckout)
		assert.Equal(t, expectedResponse, response)
	})

	t.Run("Should return a response with products applied discount and total of amount, discount and applied discount", func(t *testing.T) {
		t.Parallel()
		ctrl := gomock.NewController(t)

		mockProductCheckoutRepository := mock2.NewMockProductCheckoutRepository(ctrl)
		mockDiscountServiceRepository := mock2.NewMockDiscountServiceRepository(ctrl)

		gomock.InOrder(
			mockDiscountServiceRepository.EXPECT().GetProductDiscount(gomock.Any()).Return(0.15, nil),
			mockDiscountServiceRepository.EXPECT().GetProductDiscount(gomock.Any()).Return(0.25, nil),
		)
		mockProductCheckoutRepository.EXPECT().GetProducts(gomock.Any()).Return(mock2.ProductsToApplyDiscountResponse)

		sut := use_cases.NewProductsCheckout(mockProductCheckoutRepository, mockDiscountServiceRepository, blackFridayDate)

		response := sut.CheckoutProducts(productsCheckout)
		assert.Equal(t, mock2.CheckoutResponse, response)
	})
}

func TestApplyDiscountToProducts(t *testing.T) {
	t.Run("Should apply discount equals to 0 if service discount returns error", func(t *testing.T) {
		t.Parallel()
		ctrl := gomock.NewController(t)

		mockProductCheckoutRepository := mock2.NewMockProductCheckoutRepository(ctrl)
		mockDiscountServiceRepository := mock2.NewMockDiscountServiceRepository(ctrl)

		gomock.InOrder(
			mockDiscountServiceRepository.EXPECT().GetProductDiscount(gomock.Any()).Return(0.0, errors.New("service unavailable")),
			mockDiscountServiceRepository.EXPECT().GetProductDiscount(gomock.Any()).Return(0.25, nil),
		)

		sut := use_cases.NewProductsCheckout(mockProductCheckoutRepository, mockDiscountServiceRepository, blackFridayDate)
		productsAppliedDiscount := use_cases.ApplyDiscountToProducts(sut, mock2.ProductsToApplyDiscountResponse)

		assert.Equal(t, mock2.ProductsNotAppliedDiscount, productsAppliedDiscount)
	})

	t.Run("Should apply discount equals all products if service return success", func(t *testing.T) {
		t.Parallel()
		ctrl := gomock.NewController(t)

		mockProductCheckoutRepository := mock2.NewMockProductCheckoutRepository(ctrl)
		mockDiscountServiceRepository := mock2.NewMockDiscountServiceRepository(ctrl)

		gomock.InOrder(
			mockDiscountServiceRepository.EXPECT().GetProductDiscount(gomock.Any()).Return(0.15, nil),
			mockDiscountServiceRepository.EXPECT().GetProductDiscount(gomock.Any()).Return(0.15, nil),
		)

		sut := use_cases.NewProductsCheckout(mockProductCheckoutRepository, mockDiscountServiceRepository, blackFridayDate)
		productsAppliedDiscount := use_cases.ApplyDiscountToProducts(sut, mock2.ProductsToApplyDiscountWithoutGift)

		assert.Equal(t, mock2.ProductsAppliedDiscountWithoutGift, productsAppliedDiscount)
	})
}

func TestSumTotalForResponse(t *testing.T) {
	t.Run("Should sum all products and return by reference", func(t *testing.T) {
		t.Parallel()
		var totalAmount int64
		var totalDiscount int64
		var totalAmountWithDiscount int64

		var expectedTotalAmount int64
		var expectedTotalDiscount int64

		for _, product := range mock2.ProductsAppliedDiscountWithoutGift {
			expectedTotalAmount += product.TotalAmount
			expectedTotalDiscount += product.Discount
		}

		expectedTotalAmountWithDiscount := expectedTotalAmount - expectedTotalDiscount

		use_cases.SumTotalForResponse(mock2.ProductsAppliedDiscountWithoutGift, &totalAmount, &totalDiscount, &totalAmountWithDiscount)

		assert.Equal(t, expectedTotalAmount, totalAmount)
		assert.Equal(t, expectedTotalDiscount, totalDiscount)
		assert.Equal(t, expectedTotalAmountWithDiscount, totalAmountWithDiscount)
	})
}

func TestAddGiftToCheckout(t *testing.T) {
	t.Run("Should add a gift in product list", func(t *testing.T) {
		t.Parallel()
		ctrl := gomock.NewController(t)

		mockProductCheckoutRepository := mock2.NewMockProductCheckoutRepository(ctrl)
		mockDiscountServiceRepository := mock2.NewMockDiscountServiceRepository(ctrl)

		giftProduct := protocols.ProductAppliedDiscount{
			ID:          mock2.Product3.ID,
			Quantity:    2,
			UnitAmount:  0,
			TotalAmount: 0,
			Discount:    0,
			IsGift:      mock2.Product3.IsGift,
		}
		productsAppliedDiscountWithoutGift := mock2.ProductsAppliedDiscountWithoutGift

		mockProductCheckoutRepository.EXPECT().GetProductToGift().Return(giftProduct)

		sut := use_cases.NewProductsCheckout(mockProductCheckoutRepository, mockDiscountServiceRepository, time.Now())

		use_cases.AddGiftToCheckout(sut, &productsAppliedDiscountWithoutGift)


		assert.Equal(t, mock2.AllProductsAppliedDiscountWithGift, productsAppliedDiscountWithoutGift)

	})
}

func TestExistsGiftAddedInProducts(t *testing.T) {
	t.Run("Should return false if not exists product gift in list", func(t *testing.T) {
		t.Parallel()
		expectedResponse := false

		existsGiftInList := use_cases.ExistsGiftAddedInProducts(mock2.ProductsAppliedDiscountWithoutGift)
		assert.Equal(t, expectedResponse, existsGiftInList)
	})

	t.Run("Should return true if exists product gift in list", func(t *testing.T) {
		t.Parallel()
		expectedResponse := true

		existsGiftInList := use_cases.ExistsGiftAddedInProducts(mock2.ProductsAppliedDiscount)
		assert.Equal(t, expectedResponse, existsGiftInList)
	})
}
