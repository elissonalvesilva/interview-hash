package protocols

import valueObjects "github.com/elissonalvesilva/interview-hash/my-ecommerce/domain/value-objects"

type ProductAppliedDiscount struct {
	ID valueObjects.ID `json:"id"`
	Quantity int `json:"quantity"`
	UnitAmount float64 `json:"unit_amount"`
	TotalAmount float64 `json:"total_amount"`
	Discount float64 `json:"discount"`
	IsGift bool `json:"is_gift"`
}
