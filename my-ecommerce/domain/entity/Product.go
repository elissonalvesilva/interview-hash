package entity

import valueObjects "github.com/elissonalvesilva/interview-hash/my-ecommerce/domain/value-objects"

type Product struct {
	ID          valueObjects.ID `json:"id"`
	Title       valueObjects.Title `json:"title"`
	Description valueObjects.Description `json:"description"`
	Amount      int    `json:"amount"`
	IsGift      bool   `json:"is_gift"`
}
