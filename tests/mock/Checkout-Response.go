package mock

import (
	"github.com/elissonalvesilva/interview-hash/my-ecommerce/domain/protocols"
	"github.com/elissonalvesilva/interview-hash/my-ecommerce/shared/currency"
)

var defaultDiscount = 0.15
var quantityProduct1 = 2
var quantityProduct2 = 2

var ProductsToCheckout = []protocols.ProductCheckout{
	{ID: 1, Quantity: 2},
	{ID: 2, Quantity: 2},
	{ID: 3, Quantity: 2},
}

var ProductsAppliedDiscount = []protocols.ProductAppliedDiscount {
	protocols.ProductAppliedDiscount{
		ID:          Product1.ID,
		Quantity:    quantityProduct1,
		UnitAmount:  currency.ParseToCents(currency.TruncateNaive(Product1.Amount)),
		TotalAmount: currency.ParseToCents(currency.TruncateNaive(Product1.Amount * float64(quantityProduct1))),
		Discount:    currency.ParseToCents(currency.TruncateNaive(Product1.Amount * float64(quantityProduct1) * defaultDiscount)),
		IsGift:      Product1.IsGift,
	},
	protocols.ProductAppliedDiscount{
		ID:          Product3.ID,
		Quantity:    quantityProduct2,
		UnitAmount:  0.0,
		TotalAmount: 0.0,
		Discount:    0.0,
		IsGift:      Product3.IsGift,
	},
}

var ProductsNotAppliedDiscount = []protocols.ProductAppliedDiscount {
	protocols.ProductAppliedDiscount{
		ID:          Product1.ID,
		Quantity:    quantityProduct1,
		UnitAmount:  currency.ParseToCents(currency.TruncateNaive(Product1.Amount)),
		TotalAmount: currency.ParseToCents(currency.TruncateNaive(Product1.Amount * float64(quantityProduct1))),
		Discount:    0,
		IsGift:      Product1.IsGift,
	},
	protocols.ProductAppliedDiscount{
		ID:          Product3.ID,
		Quantity:    quantityProduct2,
		UnitAmount:  0.0,
		TotalAmount: 0.0,
		Discount:    0.0,
		IsGift:      Product3.IsGift,
	},
}

var ProductsAppliedDiscountWithoutGift = []protocols.ProductAppliedDiscount {
	protocols.ProductAppliedDiscount{
		ID:          Product1.ID,
		Quantity:    quantityProduct1,
		UnitAmount:  currency.ParseToCents(currency.TruncateNaive(Product1.Amount)),
		TotalAmount: currency.ParseToCents(currency.TruncateNaive(Product1.Amount * float64(quantityProduct1))),
		Discount:    currency.ParseToCents(currency.TruncateNaive(Product1.Amount * float64(quantityProduct1) * defaultDiscount)),
		IsGift:      Product1.IsGift,
	},
	protocols.ProductAppliedDiscount{
		ID:          Product2.ID,
		Quantity:    quantityProduct2,
		UnitAmount:  currency.ParseToCents(currency.TruncateNaive(Product2.Amount)),
		TotalAmount: currency.ParseToCents(currency.TruncateNaive(Product2.Amount * float64(quantityProduct1))),
		Discount:    currency.ParseToCents(currency.TruncateNaive(Product2.Amount * float64(quantityProduct1) * defaultDiscount)),
		IsGift:      Product2.IsGift,
	},
}

var AllProductsAppliedDiscountWithGift = []protocols.ProductAppliedDiscount {
	protocols.ProductAppliedDiscount{
		ID:          Product1.ID,
		Quantity:    quantityProduct1,
		UnitAmount:  currency.ParseToCents(currency.TruncateNaive(Product1.Amount)),
		TotalAmount: currency.ParseToCents(currency.TruncateNaive(Product1.Amount * float64(quantityProduct1))),
		Discount:    currency.ParseToCents(currency.TruncateNaive(Product1.Amount * float64(quantityProduct1) * defaultDiscount)),
		IsGift:      Product1.IsGift,
	},
	protocols.ProductAppliedDiscount{
		ID:          Product2.ID,
		Quantity:    quantityProduct2,
		UnitAmount:  currency.ParseToCents(currency.TruncateNaive(Product2.Amount)),
		TotalAmount: currency.ParseToCents(currency.TruncateNaive(Product2.Amount * float64(quantityProduct1))),
		Discount:    currency.ParseToCents(currency.TruncateNaive(Product2.Amount * float64(quantityProduct1) * defaultDiscount)),
		IsGift:      Product2.IsGift,
	},
	protocols.ProductAppliedDiscount{
		ID:          Product3.ID,
		Quantity:    quantityProduct2,
		UnitAmount:  0,
		TotalAmount: 0,
		Discount:    0,
		IsGift:      Product3.IsGift,
	},
}

var ProductsToApplyDiscountResponse = []protocols.ProductToApplyDiscount {
	protocols.ProductToApplyDiscount{
		ID:       Product1.ID,
		Amount:   Product1.Amount,
		Quantity: quantityProduct1,
		IsGift:   Product1.IsGift,
	},
	protocols.ProductToApplyDiscount{
		ID:       Product3.ID,
		Amount:   Product3.Amount,
		Quantity: quantityProduct2,
		IsGift:   Product3.IsGift,
	},
}

var AllProductsToApplyDiscount = []protocols.ProductToApplyDiscount {
	protocols.ProductToApplyDiscount{
		ID:       Product1.ID,
		Amount:   Product1.Amount,
		Quantity: quantityProduct1,
		IsGift:   Product1.IsGift,
	},
	protocols.ProductToApplyDiscount{
		ID:       Product2.ID,
		Amount:   Product2.Amount,
		Quantity: quantityProduct2,
		IsGift:   Product2.IsGift,
	},
	protocols.ProductToApplyDiscount{
		ID:       Product3.ID,
		Amount:   Product3.Amount,
		Quantity: quantityProduct2,
		IsGift:   Product3.IsGift,
	},
}

var ProductsToApplyDiscountWithoutGift = []protocols.ProductToApplyDiscount {
	protocols.ProductToApplyDiscount{
		ID:       Product1.ID,
		Amount:   Product1.Amount,
		Quantity: quantityProduct1,
		IsGift:   Product1.IsGift,
	},
	protocols.ProductToApplyDiscount{
		ID:       Product2.ID,
		Amount:   Product2.Amount,
		Quantity: quantityProduct2,
		IsGift:   Product2.IsGift,
	},
}

var CheckoutResponse = protocols.CheckoutResponse{
	TotalAmount:             currency.ParseToCents(currency.TruncateNaive(Product1.Amount * float64(quantityProduct1))),
	TotalAmountWithDiscount: currency.ParseToCents((Product1.Amount * float64(quantityProduct1)) - (Product1.Amount * float64(quantityProduct1)) *defaultDiscount),
	TotalDiscount:           currency.ParseToCents(currency.TruncateNaive((Product1.Amount * float64(quantityProduct1)) * defaultDiscount)),
	Products:                ProductsAppliedDiscount,
}

var CheckoutResponseWithoutDiscount = protocols.CheckoutResponse{
	TotalAmount:             currency.ParseToCents(currency.TruncateNaive(Product1.Amount * float64(quantityProduct1))),
	TotalAmountWithDiscount: currency.ParseToCents(currency.TruncateNaive(Product1.Amount * float64(quantityProduct1))),
	TotalDiscount:           0,
	Products:                ProductsNotAppliedDiscount,
}

var CheckoutResponseWithGift = protocols.CheckoutResponse{
	TotalAmount:             currency.ParseToCents(currency.TruncateNaive(Product1.Amount * float64(quantityProduct1)) + currency.TruncateNaive(Product2.Amount * float64(quantityProduct2))),
	TotalAmountWithDiscount: currency.ParseToCents((Product1.Amount * float64(quantityProduct1)) - (Product1.Amount * float64(quantityProduct1)) *defaultDiscount + (Product2.Amount * float64(quantityProduct1)) - (Product2.Amount * float64(quantityProduct2)) *defaultDiscount),
	TotalDiscount:           currency.ParseToCents(currency.TruncateNaive((Product1.Amount * float64(quantityProduct1)) *defaultDiscount) + currency.TruncateNaive((Product2.Amount * float64(quantityProduct2)) *defaultDiscount)),
	Products:                ProductsAppliedDiscountWithoutGift,
}
