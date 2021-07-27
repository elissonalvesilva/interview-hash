package controllers

import (
	useCases "github.com/elissonalvesilva/interview-hash/my-ecommerce/cmd/server/factories/use-cases"
	inMemoryDB "github.com/elissonalvesilva/interview-hash/my-ecommerce/infrastructure/db/in-memory"
	"github.com/elissonalvesilva/interview-hash/my-ecommerce/infrastructure/validator/govalidator"
	"github.com/elissonalvesilva/interview-hash/my-ecommerce/presenters/controllers"
	"time"
)

func MakeCheckoutController(db *inMemoryDB.InMemoryDatabase, blackFridayDate time.Time) *controllers.CheckoutController {
	validator := govalidator.NewValidator()
	useCaseCheckout := useCases.MakeProductCheckoutUseCase(db, blackFridayDate)
	controller := controllers.NewCheckoutController(*useCaseCheckout, validator)
	return controller
}
