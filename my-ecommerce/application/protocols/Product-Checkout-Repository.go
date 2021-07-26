package protocols

import (
	"github.com/elissonalvesilva/interview-hash/my-ecommerce/domain/protocols"
)

type ProductCheckoutRepository interface {
	GetProducts([]protocols.ProductCheckout) []protocols.ProductToApplyDiscount
	GetProductToGift() protocols.ProductAppliedDiscount
}
