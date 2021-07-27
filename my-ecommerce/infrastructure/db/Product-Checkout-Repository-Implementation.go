package db

import (
	"github.com/elissonalvesilva/interview-hash/my-ecommerce/domain/protocols"
	infraProtocols "github.com/elissonalvesilva/interview-hash/my-ecommerce/infrastructure/db/in-memory/protocols"
)

type ProductCheckoutRepositoryImplementation struct {
	inMemoryDb infraProtocols.InMemory
}

func NewProductCheckoutRepositoryImplementation(db infraProtocols.InMemory) *ProductCheckoutRepositoryImplementation {
	return &ProductCheckoutRepositoryImplementation{
		inMemoryDb: db,
	}
}

func (db *ProductCheckoutRepositoryImplementation) GetProducts(products []protocols.ProductCheckout) []protocols.ProductToApplyDiscount {
	var productList []protocols.ProductToApplyDiscount

	for _, product := range products {
		productItem, errNotFound := db.inMemoryDb.GetByID(product.ID)

		if errNotFound == nil {
			item := protocols.ProductToApplyDiscount{
				ID: productItem.ID,
				Amount: productItem.Amount,
				Quantity: product.Quantity,
				IsGift: productItem.IsGift,
			}
			productList = append(productList, item)
		}
	}

	return productList
}

func (db *ProductCheckoutRepositoryImplementation) GetProductToGift() protocols.ProductAppliedDiscount {

	productGift, errGetGiftFromDb := db.inMemoryDb.GetOne(infraProtocols.Filter{Condition: "IsGift", ValueToFilter: true})

	if errGetGiftFromDb != nil {
		return protocols.ProductAppliedDiscount{}
	}

	return protocols.ProductAppliedDiscount{
		ID: productGift.ID,
		Quantity: 1,
		TotalAmount: 0,
		UnitAmount: 0,
		Discount: 0,
		IsGift: productGift.IsGift,
	}
}