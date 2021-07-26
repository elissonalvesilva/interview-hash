package protocols

type DiscountServiceRepository interface {
	GetProductDiscount(productID int) (float64, error)
}
