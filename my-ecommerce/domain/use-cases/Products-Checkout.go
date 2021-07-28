package use_cases

import domainProtocol "github.com/elissonalvesilva/interview-hash/my-ecommerce/domain/protocols"

type ProductsCheckout interface {
	CheckoutProducts(productList []domainProtocol.ProductCheckout) domainProtocol.CheckoutResponse
}