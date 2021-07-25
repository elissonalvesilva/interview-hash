package entity

import (
	"github.com/elissonalvesilva/interview-hash/my-ecommerce/domain/protocols"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApplyDiscount(t *testing.T) {
	t.Run("Should return a product as gift without prices(unit, total) values", func(t *testing.T) {
		t.Parallel()

		product := Product{
			ID: 1,
			Title: "Ergonomic Wooden Pants",
			Description: "Deleniti beatae porro.",
			Amount: 15157,
			IsGift: true,
		}

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
}
