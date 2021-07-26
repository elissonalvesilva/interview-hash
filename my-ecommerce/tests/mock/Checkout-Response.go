package mock

import (
	"github.com/elissonalvesilva/interview-hash/my-ecommerce/domain/protocols"
	"github.com/elissonalvesilva/interview-hash/my-ecommerce/shared/currency"
)

var defaultDiscount = 0.15
var quantityProduct1 = 2
var quantityProduct2 = 2

var ProductsAppliedDiscount = []protocols.ProductAppliedDiscount {
	protocols.ProductAppliedDiscount{
		ID: Product1.ID,
		Quantity: quantityProduct1,
		UnitAmount: currency.TruncateNaive(Product1.Amount),
		TotalAmount: currency.TruncateNaive(Product1.Amount * float64(quantityProduct1)),
		Discount: currency.TruncateNaive(Product1.Amount * float64(quantityProduct1) * defaultDiscount),
		IsGift: Product1.IsGift,
	},
	protocols.ProductAppliedDiscount{
		ID: Product3.ID,
		Quantity: quantityProduct2,
		UnitAmount: 0.0,
		TotalAmount: 0.0,
		Discount: 0.0,
		IsGift: Product3.IsGift,
	},
}

var ProductsToApplyDiscountResponse = []protocols.ProductToApplyDiscount {
	protocols.ProductToApplyDiscount{
		ID: Product1.ID,
		Amount: Product1.Amount,
		Quantity: quantityProduct1,
		IsGift: Product1.IsGift,
	},
	protocols.ProductToApplyDiscount{
		ID: Product3.ID,
		Amount: Product3.Amount,
		Quantity: quantityProduct2,
		IsGift: Product3.IsGift,
	},
}

var CheckoutResponse = protocols.CheckoutResponse{
	TotalAmount: currency.TruncateNaive(Product1.Amount * float64(quantityProduct1)),
	TotalAmountWithDiscount: (Product1.Amount * float64(quantityProduct1)) - (Product1.Amount * float64(quantityProduct1)) * defaultDiscount,
	TotalDiscount: currency.TruncateNaive((Product1.Amount * float64(quantityProduct1)) * defaultDiscount),
	Products: ProductsAppliedDiscount,
}
