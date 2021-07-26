package db

import "github.com/elissonalvesilva/interview-hash/my-ecommerce/domain/protocols"

type ProductCheckoutRepositoryImplementation struct {
	inMemoryDb InMemory
}

func NewProductCheckoutRepositoryImplementation(db InMemory) *ProductCheckoutRepositoryImplementation {
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
	return protocols.ProductAppliedDiscount{}
}