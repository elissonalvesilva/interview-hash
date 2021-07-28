package entity

import (
	"github.com/elissonalvesilva/interview-hash/my-ecommerce/domain/entity"
	"github.com/elissonalvesilva/interview-hash/my-ecommerce/domain/protocols"
	"github.com/elissonalvesilva/interview-hash/my-ecommerce/shared/currency"
	"github.com/stretchr/testify/assert"
	"testing"
)


var productMock = protocols.ProductToApplyDiscount{
	ID: 1,
	Amount: 15157,
	IsGift: false,
	Quantity: 1,
}

func TestApplyDiscount(t *testing.T) {
	t.Run("Should return a product as gift without prices(unit, total) values", func(t *testing.T) {
		t.Parallel()

		product := productMock
		product.IsGift = true

		quantity := 2
		discountPercentage := 0.15

		expectedProductAppliedDiscount := protocols.ProductAppliedDiscount{
			ID: product.ID,
			Quantity: quantity,
			UnitAmount: 0,
			TotalAmount: 0,
			Discount: 0,
			IsGift: product.IsGift,
		}

		appliedDiscount := entity.ApplyDiscount(product, quantity, discountPercentage)
		assert.Equal(t, expectedProductAppliedDiscount, appliedDiscount)
	})

	t.Run("Should return a product with discount not applied 'cause percentage is equals to 0", func(t *testing.T) {
		t.Parallel()

		product := productMock
		quantity := 2
		discountPercentage := 0.0

		expectedProductNotAppliedDiscount := protocols.ProductAppliedDiscount{
			ID: product.ID,
			Quantity: quantity,
			UnitAmount: currency.TruncateNaive(product.Amount),
			TotalAmount: currency.TruncateNaive(product.Amount * float64(quantity)),
			Discount: 0,
			IsGift: product.IsGift,
		}

		appliedDiscount := entity.ApplyDiscount(product, quantity, discountPercentage)
		assert.Equal(t, expectedProductNotAppliedDiscount, appliedDiscount)
	})

	t.Run("Should return a product applied discount", func(t *testing.T) {
		t.Parallel()

		product := productMock
		quantity := 2
		discountPercentage := 0.15

		expectedProductAppliedDiscount := protocols.ProductAppliedDiscount{
			ID: product.ID,
			Quantity: quantity,
			UnitAmount: currency.TruncateNaive(product.Amount),
			TotalAmount: currency.TruncateNaive(product.Amount * float64(quantity)),
			Discount: currency.TruncateNaive((product.Amount * float64(quantity)) * discountPercentage),
			IsGift: product.IsGift,
		}

		appliedDiscount := entity.ApplyDiscount(product, quantity, discountPercentage)
		assert.Equal(t, expectedProductAppliedDiscount, appliedDiscount)
	})
}
