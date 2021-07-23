package entity

import (
	valueObjects "github.com/elissonalvesilva/interview-hash/my-ecommerce/domain/value-objects"
	"github.com/elissonalvesilva/interview-hash/my-ecommerce/shared/constants"
	"time"
)

type Product struct {
	ID          valueObjects.ID `json:"id"`
	Title       valueObjects.Title `json:"title"`
	Description valueObjects.Description `json:"description"`
	Amount      int    `json:"amount"`
	IsGift      bool   `json:"is_gift"`
}

const (
	GIFT = true
)

func ApplyBlackFridayGift(product Product, dateToCompare time.Time) Product {
	if constants.BLACKFRIDAYDATE == dateToCompare {
		product.IsGift = GIFT
	}

	return product
}