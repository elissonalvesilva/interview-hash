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
	var totalAmountWithDiscount float64

	productsAppliedDiscount := ApplyDiscountToProducts(useCase, products)

	AddGiftToCheckout(useCase, &productsAppliedDiscount)

	SumTotalForResponse(productsAppliedDiscount, &totalAmount, &totalDiscount, &totalAmountWithDiscount)

	return domainProtocol.CheckoutResponse{
		TotalAmount: totalAmount,
		TotalAmountWithDiscount: totalAmountWithDiscount,
		TotalDiscount: totalDiscount,
		Products: productsAppliedDiscount,
	}
}

func ApplyDiscountToProducts(useCase *ProductCheckoutUseCase, products []domainProtocol.ProductToApplyDiscount) []domainProtocol.ProductAppliedDiscount {
	var productsAppliedDiscount []domainProtocol.ProductAppliedDiscount

	for _, product := range products {
		discount, serviceError := useCase.service.GetProductDiscount(int(product.ID))
		if serviceError != nil {
			discount = 0
		}

		productAppliedDiscount := entity.ApplyDiscount(product, product.Quantity, discount)
		productsAppliedDiscount = append(productsAppliedDiscount, productAppliedDiscount)
	}

	return productsAppliedDiscount
}

func SumTotalForResponse(productsAppliedDiscount []domainProtocol.ProductAppliedDiscount, totalAmount *float64, totalDiscount *float64, totalAmountWithDiscount *float64) {
	for _, product := range productsAppliedDiscount {
		*totalAmount += product.TotalAmount
		*totalDiscount += product.Discount
	}

	*totalAmountWithDiscount = *totalAmount - *totalDiscount
}

func AddGiftToCheckout(useCase *ProductCheckoutUseCase, productsAppliedDiscount *[]domainProtocol.ProductAppliedDiscount) {
	if date.IsBlackFriday(time.Now(), useCase.blackFridayDate) {
		if !ExistsGiftAddedInProducts(*productsAppliedDiscount) {
			productGift := useCase.repo.GetProductToGift()
			*productsAppliedDiscount = append(*productsAppliedDiscount, productGift)
		}
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