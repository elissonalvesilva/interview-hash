package use_cases

import (
	useCaseProtocol "github.com/elissonalvesilva/interview-hash/my-ecommerce/application/protocols"
	domainProtocol "github.com/elissonalvesilva/interview-hash/my-ecommerce/domain/protocols"
)

type ProductCheckoutUseCase struct {
	repo useCaseProtocol.ProductCheckoutRepository
}

func NewProductsCheckout(repository useCaseProtocol.ProductCheckoutRepository) *ProductCheckoutUseCase {
	return &ProductCheckoutUseCase{
		repo: repository,
	}
}

func (useCase *ProductCheckoutUseCase) CheckoutProducts(productList []domainProtocol.ProductCheckout) domainProtocol.CheckoutResponse {

	products := useCase.repo.GetProducts(productList)

	if len(products) == 0 {
		return domainProtocol.CheckoutResponse{}
	}


	return domainProtocol.CheckoutResponse{}
}