package controllers

import (
	"encoding/json"
	useCases "github.com/elissonalvesilva/interview-hash/my-ecommerce/application/use-cases"
	domainProtocol "github.com/elissonalvesilva/interview-hash/my-ecommerce/domain/protocols"
	errorToResponse "github.com/elissonalvesilva/interview-hash/my-ecommerce/presenters/error"
	sendResponse "github.com/elissonalvesilva/interview-hash/my-ecommerce/presenters/helpers"
	presentersProtocol "github.com/elissonalvesilva/interview-hash/my-ecommerce/presenters/protocols"
	"net/http"
)

type CheckoutController struct {
	useCase useCases.ProductCheckoutUseCase
	validator presentersProtocol.ValidateParam
}

func NewCheckoutController(useCase useCases.ProductCheckoutUseCase, validator presentersProtocol.ValidateParam) *CheckoutController {
	return &CheckoutController{
		useCase: useCase,
		validator: validator,
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

	errorValidationParam := ctrl.validator.ValidateRequestParams(productCheckoutRequest)
	if errorValidationParam != nil {
		sendResponse.BadRequest(w, errorToResponse.InvalidRequestParams(errorValidationParam))
		return
	}

	checkoutProduct := ctrl.useCase.CheckoutProducts(productCheckoutRequest.Products)
	if len(checkoutProduct.Products) == 0 {
		var ids []int

		for _, product := range checkoutProduct.Products {
			ids = append(ids, int(product.ID))
		}

		sendResponse.NotFound(w, errorToResponse.NotFoundProducts(ids))
		return
	}


	sendResponse.Ok(w, checkoutProduct)
	return
}
