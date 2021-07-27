package controllers

import (
	"encoding/json"
	"github.com/asaskevich/govalidator"
	useCases "github.com/elissonalvesilva/interview-hash/my-ecommerce/application/use-cases"
	domainProtocol "github.com/elissonalvesilva/interview-hash/my-ecommerce/domain/protocols"
	"net/http"
)

type CheckoutController struct {
	useCase useCases.ProductCheckoutUseCase
}

func NewCheckoutController(useCase useCases.ProductCheckoutUseCase) *CheckoutController {
	return &CheckoutController{
		useCase: useCase,
	}
}

func (ctrl *CheckoutController) CheckoutProductsController (w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var productCheckoutRequest domainProtocol.ProductCheckoutRequest
	errDecode := decoder.Decode(&productCheckoutRequest)

	if errDecode != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid request param json")
		return
	}

	_, errorValidationParam := govalidator.ValidateStruct(productCheckoutRequest)
	if errorValidationParam != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorValidationParam)
		return
	}

	checkoutProduct := ctrl.useCase.CheckoutProducts(productCheckoutRequest.Products)

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(checkoutProduct)
	return
}
