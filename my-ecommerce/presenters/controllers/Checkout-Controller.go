package controllers

import (
	"encoding/json"
	"github.com/asaskevich/govalidator"
	useCases "github.com/elissonalvesilva/interview-hash/my-ecommerce/application/use-cases"
	domainProtocol "github.com/elissonalvesilva/interview-hash/my-ecommerce/domain/protocols"
	errorToResponse "github.com/elissonalvesilva/interview-hash/my-ecommerce/presenters/error"
	sendResponse "github.com/elissonalvesilva/interview-hash/my-ecommerce/presenters/helpers"
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
		sendResponse.BadRequest(w, errorToResponse.InvalidJsonParamResponse())
		return
	}

	_, errorValidationParam := govalidator.ValidateStruct(productCheckoutRequest)
	if errorValidationParam != nil {
		sendResponse.BadRequest(w, errorToResponse.InvalidRequestParams(errorValidationParam))
		return
	}

	checkoutProduct := ctrl.useCase.CheckoutProducts(productCheckoutRequest.Products)

	sendResponse.Ok(w, checkoutProduct)
	return
}
