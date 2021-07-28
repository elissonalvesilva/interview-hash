package db

import (
	"github.com/elissonalvesilva/interview-hash/my-ecommerce/domain/entity"
	"github.com/elissonalvesilva/interview-hash/my-ecommerce/domain/protocols"
	"github.com/elissonalvesilva/interview-hash/my-ecommerce/infrastructure/db"
	inMemoryDB "github.com/elissonalvesilva/interview-hash/my-ecommerce/infrastructure/db/in-memory"
	mock2 "github.com/elissonalvesilva/interview-hash/tests/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductCheckoutRepositoryImplementation_GetProducts(t *testing.T) {
	t.Run("Should return a empty list when not found any product in db", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		mockInMemoryDatabase := mock2.NewMockInMemory(ctrl)
		productListToApplyDiscount := mock2.ProductsToCheckout

		gomock.InOrder(
			mockInMemoryDatabase.EXPECT().GetByID(1).Return(entity.Product{}, inMemoryDB.ErrNotFoundItemInDB),
			mockInMemoryDatabase.EXPECT().GetByID(2).Return(entity.Product{}, inMemoryDB.ErrNotFoundItemInDB),
			mockInMemoryDatabase.EXPECT().GetByID(3).Return(entity.Product{}, inMemoryDB.ErrNotFoundItemInDB),
		)

		sut := db.NewProductCheckoutRepositoryImplementation(mockInMemoryDatabase)

		response := sut.GetProducts(productListToApplyDiscount)
		assert.Equal(t, []protocols.ProductToApplyDiscount(nil), response)
	})

	t.Run("Should return a list of products to apply discount", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		mockInMemoryDatabase := mock2.NewMockInMemory(ctrl)
		productListToApplyDiscount := mock2.ProductsToCheckout

		gomock.InOrder(
			mockInMemoryDatabase.EXPECT().GetByID(1).Return(mock2.Product1, nil),
			mockInMemoryDatabase.EXPECT().GetByID(2).Return(mock2.Product2, nil),
			mockInMemoryDatabase.EXPECT().GetByID(3).Return(mock2.Product3, nil),
		)

		sut := db.NewProductCheckoutRepositoryImplementation(mockInMemoryDatabase)

		response := sut.GetProducts(productListToApplyDiscount)
		assert.Equal(t, mock2.AllProductsToApplyDiscount, response)
	})
}

func TestProductCheckoutRepositoryImplementation_GetProductToGift(t *testing.T) {
	t.Run("Should return GetOne throws or not found filter condition", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		mockInMemoryDatabase := mock2.NewMockInMemory(ctrl)

		mockInMemoryDatabase.EXPECT().GetOne(gomock.Any()).Return(entity.Product{}, inMemoryDB.ErrNotFoundItemInDBByCondition)

		sut := db.NewProductCheckoutRepositoryImplementation(mockInMemoryDatabase)

		response := sut.GetProductToGift()
		assert.Equal(t, protocols.ProductAppliedDiscount{}, response)
	})

	t.Run("Should return product gift to added", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		mockInMemoryDatabase := mock2.NewMockInMemory(ctrl)

		mockInMemoryDatabase.EXPECT().GetOne(gomock.Any()).Return(mock2.Product3, nil)

		sut := db.NewProductCheckoutRepositoryImplementation(mockInMemoryDatabase)

		response := sut.GetProductToGift()
		assert.Equal(t, protocols.ProductAppliedDiscount{
			ID:          mock2.Product3.ID,
			Quantity:    1,
			TotalAmount: 0,
			UnitAmount:  0,
			Discount:    0,
			IsGift:      mock2.Product3.IsGift,
		}, response)
	})
}
