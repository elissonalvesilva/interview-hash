package protocols

type CheckoutResponse struct {
	TotalAmount int64 `json:"total_amount"`
	TotalAmountWithDiscount int64 `json:"total_amount_with_discount"`
	TotalDiscount int64 `json:"total_discount"`
	Products []ProductAppliedDiscount `json:"products"`
}
