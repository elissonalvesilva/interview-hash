package protocols

type CheckoutResponse struct {
	TotalAmount float64 `json:"total_amount"`
	TotalAmountWithDiscount float64 `json:"total_amount_with_discount"`
	TotalDiscount float64 `json:"total_discount"`
	Products []ProductAppliedDiscount `json:"products"`
}
