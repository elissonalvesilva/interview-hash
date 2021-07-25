package entity

import (
	"github.com/elissonalvesilva/interview-hash/my-ecommerce/domain/protocols"
	"github.com/stretchr/testify/assert"
	"testing"
)


var productMock = Product{
	ID: 1,
	Title: "Ergonomic Wooden Pants",
	Description: "Deleniti beatae porro.",
	Amount: 15157,
	IsGift: false,
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

		appliedDiscount := ApplyDiscount(product, quantity, discountPercentage)
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
			UnitAmount: product.Amount,
			TotalAmount: product.Amount * float64(quantity),
			Discount: 0,
			IsGift: product.IsGift,
		}

		appliedDiscount := ApplyDiscount(product, quantity, discountPercentage)
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
			UnitAmount: product.Amount,
			TotalAmount: product.Amount * float64(quantity),
			Discount: product.Amount * discountPercentage,
			IsGift: product.IsGift,
		}

		appliedDiscount := ApplyDiscount(product, quantity, discountPercentage)
		assert.Equal(t, expectedProductAppliedDiscount, appliedDiscount)
	})
}
