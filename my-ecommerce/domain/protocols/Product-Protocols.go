package protocols

import valueObjects "github.com/elissonalvesilva/interview-hash/my-ecommerce/domain/value-objects"

type ProductCheckout struct {
	ID int
	Quantity int
}

type ProductToApplyDiscount struct {
	ID valueObjects.ID
	Amount float64
	Quantity int
	IsGift bool
}
