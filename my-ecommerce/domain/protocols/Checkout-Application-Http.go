package protocols

type ProductCheckoutRequest struct {
	Products []ProductCheckout `json:"products" valid:"required"`
}
