package use_cases

import (
	useCaseProtocol "github.com/elissonalvesilva/interview-hash/my-ecommerce/application/protocols"
	"github.com/elissonalvesilva/interview-hash/my-ecommerce/domain/entity"
	domainProtocol "github.com/elissonalvesilva/interview-hash/my-ecommerce/domain/protocols"
)

type ProductCheckoutUseCase struct {
	repo useCaseProtocol.ProductCheckoutRepository
	service useCaseProtocol.DiscountServiceRepository
}

func NewProductsCheckout(repository useCaseProtocol.ProductCheckoutRepository, service useCaseProtocol.DiscountServiceRepository) *ProductCheckoutUseCase {
	return &ProductCheckoutUseCase{
		repo: repository,
		service: service,
	}
}

func (useCase *ProductCheckoutUseCase) CheckoutProducts(productList []domainProtocol.ProductCheckout) domainProtocol.CheckoutResponse {

	products := useCase.repo.GetProducts(productList)

	if len(products) == 0 {
		return domainProtocol.CheckoutResponse{}
	}

	var totalAmount float64
	var totalDiscount float64
	var productsAppliedDiscount []domainProtocol.ProductAppliedDiscount

	for _, product := range products {
		discount, err := useCase.service.GetProductDiscount(int(product.ID))
		if err != nil {
			discount = 0
		}

		productAppliedDiscount := entity.ApplyDiscount(product, product.Quantity, discount)
		totalAmount += productAppliedDiscount.TotalAmount
		totalDiscount += productAppliedDiscount.Discount
		productsAppliedDiscount = append(productsAppliedDiscount, productAppliedDiscount)
	}

	totalAmountWithDiscount := totalAmount - totalDiscount
	return domainProtocol.CheckoutResponse{
		TotalAmount: totalAmount,
		TotalAmountWithDiscount: totalAmountWithDiscount,
		TotalDiscount: totalDiscount,
		Products: productsAppliedDiscount,
	}
}