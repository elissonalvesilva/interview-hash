package use_cases

import (
	useCaseProtocol "github.com/elissonalvesilva/interview-hash/my-ecommerce/application/protocols"
	"github.com/elissonalvesilva/interview-hash/my-ecommerce/domain/entity"
	domainProtocol "github.com/elissonalvesilva/interview-hash/my-ecommerce/domain/protocols"
	"github.com/elissonalvesilva/interview-hash/my-ecommerce/shared/date"
	"time"
)

type ProductCheckoutUseCase struct {
	repo useCaseProtocol.ProductCheckoutRepository
	service useCaseProtocol.DiscountServiceRepository
	blackFridayDate time.Time
}

func NewProductsCheckout(repository useCaseProtocol.ProductCheckoutRepository, service useCaseProtocol.DiscountServiceRepository, blackFridayDate time.Time) *ProductCheckoutUseCase {
	return &ProductCheckoutUseCase{
		repo: repository,
		service: service,
		blackFridayDate: blackFridayDate,
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
		discount, serviceError := useCase.service.GetProductDiscount(int(product.ID))
		if serviceError != nil {
			discount = 0
		}

		productAppliedDiscount := entity.ApplyDiscount(product, product.Quantity, discount)
		totalAmount += productAppliedDiscount.TotalAmount
		totalDiscount += productAppliedDiscount.Discount
		productsAppliedDiscount = append(productsAppliedDiscount, productAppliedDiscount)
	}

	if date.IsBlackFriday(time.Now(), useCase.blackFridayDate) {
		if !ExistsGiftAddedInProducts(productsAppliedDiscount) {
			productGift := useCase.repo.GetProductToGift()
			productsAppliedDiscount = append(productsAppliedDiscount, productGift)
		}
	}

	totalAmountWithDiscount := totalAmount - totalDiscount
	return domainProtocol.CheckoutResponse{
		TotalAmount: totalAmount,
		TotalAmountWithDiscount: totalAmountWithDiscount,
		TotalDiscount: totalDiscount,
		Products: productsAppliedDiscount,
	}
}

func ExistsGiftAddedInProducts(products []domainProtocol.ProductAppliedDiscount) bool {
	existsProductGift := false

	for _, product := range products {
		if product.IsGift {
			existsProductGift = true
			return existsProductGift
		}
	}

	return existsProductGift
}