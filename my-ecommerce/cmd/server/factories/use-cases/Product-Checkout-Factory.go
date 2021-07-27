package use_cases

import (
	useCases "github.com/elissonalvesilva/interview-hash/my-ecommerce/application/use-cases"
	repositoryImplementation "github.com/elissonalvesilva/interview-hash/my-ecommerce/infrastructure/db"
	inMemoryDB "github.com/elissonalvesilva/interview-hash/my-ecommerce/infrastructure/db/in-memory"
	"github.com/elissonalvesilva/interview-hash/my-ecommerce/infrastructure/grpc"
	"time"
)

func MakeProductCheckoutUseCase(db *inMemoryDB.InMemoryDatabase, blackFridayDate time.Time) *useCases.ProductCheckoutUseCase {
	productCheckoutRepository := repositoryImplementation.NewProductCheckoutRepositoryImplementation(db)
	discountServiceRepository := grpc.NewDiscountGRPCService()
	productCheckoutUseCase := useCases.NewProductsCheckout(productCheckoutRepository, discountServiceRepository ,blackFridayDate)
	return productCheckoutUseCase
}
