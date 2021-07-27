package db

import (
	"github.com/elissonalvesilva/interview-hash/my-ecommerce/domain/entity"
	"github.com/elissonalvesilva/interview-hash/my-ecommerce/domain/protocols"
	inMemoryDB "github.com/elissonalvesilva/interview-hash/my-ecommerce/infrastructure/db/in-memory"
	"github.com/elissonalvesilva/interview-hash/my-ecommerce/tests/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductCheckoutRepositoryImplementation_GetProducts(t *testing.T) {
	t.Run("Should return a empty list when not found any product in db", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		mockInMemoryDatabase := mock.NewMockInMemory(ctrl)
		productListToApplyDiscount := mock.ProductsToCheckout

		gomock.InOrder(
			mockInMemoryDatabase.EXPECT().GetByID(1).Return(entity.Product{}, inMemoryDB.ErrNotFoundItemInDB),
			mockInMemoryDatabase.EXPECT().GetByID(2).Return(entity.Product{}, inMemoryDB.ErrNotFoundItemInDB),
			mockInMemoryDatabase.EXPECT().GetByID(3).Return(entity.Product{}, inMemoryDB.ErrNotFoundItemInDB),
		)

		sut := NewProductCheckoutRepositoryImplementation(mockInMemoryDatabase)

		response := sut.GetProducts(productListToApplyDiscount)
		assert.Equal(t, []protocols.ProductToApplyDiscount(nil), response)
	})

	t.Run("Should return a list of products to apply discount", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		mockInMemoryDatabase := mock.NewMockInMemory(ctrl)
		productListToApplyDiscount := mock.ProductsToCheckout

		gomock.InOrder(
			mockInMemoryDatabase.EXPECT().GetByID(1).Return(mock.Product1, nil),
			mockInMemoryDatabase.EXPECT().GetByID(2).Return(mock.Product2, nil),
			mockInMemoryDatabase.EXPECT().GetByID(3).Return(mock.Product3, nil),
		)

		sut := NewProductCheckoutRepositoryImplementation(mockInMemoryDatabase)

		response := sut.GetProducts(productListToApplyDiscount)
		assert.Equal(t, mock.AllProductsToApplyDiscount, response)
	})
}

func TestProductCheckoutRepositoryImplementation_GetProductToGift(t *testing.T) {
	t.Run("Should return GetOne throws or not found filter condition", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		mockInMemoryDatabase := mock.NewMockInMemory(ctrl)

		mockInMemoryDatabase.EXPECT().GetOne(gomock.Any()).Return(entity.Product{}, inMemoryDB.ErrNotFoundItemInDBByCondition)

		sut := NewProductCheckoutRepositoryImplementation(mockInMemoryDatabase)

		response := sut.GetProductToGift()
		assert.Equal(t, protocols.ProductAppliedDiscount{}, response)
	})

	t.Run("Should return product gift to added", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		mockInMemoryDatabase := mock.NewMockInMemory(ctrl)

		mockInMemoryDatabase.EXPECT().GetOne(gomock.Any()).Return(mock.Product3, nil)

		sut := NewProductCheckoutRepositoryImplementation(mockInMemoryDatabase)

		response := sut.GetProductToGift()
		assert.Equal(t, protocols.ProductAppliedDiscount{
			ID: mock.Product3.ID,
			Quantity: 1,
			TotalAmount: 0,
			UnitAmount: 0,
			Discount: 0,
			IsGift: mock.Product3.IsGift,
		}, response)
	})
}
