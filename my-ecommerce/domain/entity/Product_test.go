package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var product = Product{
	ID: 1,
	Title: "Ergonomic Wooden Pants",
	Description: "Deleniti beatae porro.",
	Amount: 15157,
	IsGift: false,
}

func TestApplyBlackFridayGift(t *testing.T) {
	t.Run("Should not apply gift if is not equal to black friday date", func(t *testing.T) {
		t.Parallel()

		productBeforeFunction := product
		fakeBlackFridayDate := time.Now()

		productAfterFunction := ApplyBlackFridayGift(productBeforeFunction, fakeBlackFridayDate)
		assert.Equal(t, productBeforeFunction, productAfterFunction)
	})

	t.Run("Should not apply gift if is not equal to black friday date", func(t *testing.T) {
		t.Parallel()

		product := product

		productAfterAppliedGift := Product{
			ID: 1,
			Title: "Ergonomic Wooden Pants",
			Description: "Deleniti beatae porro.",
			Amount: 15157,
			IsGift: true,
		}
		fakeBlackFridayDate := time.Date(2021, 11, 29, 00, 00, 00, 00, time.UTC)

		productAppliedGift := ApplyBlackFridayGift(product, fakeBlackFridayDate)
		assert.Equal(t, productAfterAppliedGift, productAppliedGift)
	})
}
