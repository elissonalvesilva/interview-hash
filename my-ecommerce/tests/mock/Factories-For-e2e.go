package mock

import (
	useCases "github.com/elissonalvesilva/interview-hash/my-ecommerce/application/use-cases"
	repositoryImplementation "github.com/elissonalvesilva/interview-hash/my-ecommerce/infrastructure/db"
	inMemoryDB "github.com/elissonalvesilva/interview-hash/my-ecommerce/infrastructure/db/in-memory"
	"github.com/elissonalvesilva/interview-hash/my-ecommerce/infrastructure/validator/govalidator"
	"github.com/elissonalvesilva/interview-hash/my-ecommerce/presenters/controllers"
	"github.com/golang/mock/gomock"
	"testing"
	"time"
)

func SetupDiscountService(t *testing.T, needsDiscountService bool) *MockDiscountServiceRepository {
	ctrl := gomock.NewController(t)
	mockDiscountServiceRepository := NewMockDiscountServiceRepository(ctrl)
	if needsDiscountService {
		gomock.InOrder(
			mockDiscountServiceRepository.EXPECT().GetProductDiscount(gomock.Any()).Return(0.15, nil),
			mockDiscountServiceRepository.EXPECT().GetProductDiscount(gomock.Any()).Return(0.15, nil),
		)
	}
	return mockDiscountServiceRepository
}

func SetupController(t *testing.T, db *inMemoryDB.InMemoryDatabase, blackFridayDate time.Time, needsDiscountService bool) *controllers.CheckoutController {
	validator := govalidator.NewValidator()
	productCheckoutRepository := repositoryImplementation.NewProductCheckoutRepositoryImplementation(db)
	discountServiceRepository := SetupDiscountService(t, needsDiscountService)
	productCheckoutUseCase := useCases.NewProductsCheckout(productCheckoutRepository, discountServiceRepository,blackFridayDate)
	controller := controllers.NewCheckoutController(*productCheckoutUseCase, validator)
	return controller
}
