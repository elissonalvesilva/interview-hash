package entity

import (
	"github.com/elissonalvesilva/interview-hash/my-ecommerce/domain/protocols"
	valueObjects "github.com/elissonalvesilva/interview-hash/my-ecommerce/domain/value-objects"
)

type Product struct {
	ID          valueObjects.ID `json:"id"`
	Title       valueObjects.Title `json:"title"`
	Description valueObjects.Description `json:"description"`
	Amount      float64    `json:"amount"`
	IsGift      bool   `json:"is_gift"`
}

const ISGIFT = true

func ApplyDiscount(product Product, quantity int, percentage float64) protocols.ProductAppliedDiscount {

	if product.IsGift == ISGIFT {
		return protocols.ProductAppliedDiscount{
			ID: product.ID,
			Quantity: quantity,
			UnitAmount: 0,
			TotalAmount: 0,
			Discount: 0,
			IsGift: product.IsGift,
		}
	}

	totalAmount := product.Amount * float64(quantity)
	if percentage == 0.0 {
		return protocols.ProductAppliedDiscount{
			ID: product.ID,
			Quantity: quantity,
			UnitAmount: product.Amount,
			TotalAmount: totalAmount,
			Discount: 0,
			IsGift: product.IsGift,
		}
	}

	discount := product.Amount * percentage
	return protocols.ProductAppliedDiscount{
		ID: product.ID,
		Quantity: quantity,
		UnitAmount: product.Amount,
		TotalAmount: totalAmount,
		Discount: discount,
		IsGift: product.IsGift,
	}
}