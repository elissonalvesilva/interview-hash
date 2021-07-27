package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	useCases "github.com/elissonalvesilva/interview-hash/my-ecommerce/application/use-cases"
	"github.com/elissonalvesilva/interview-hash/my-ecommerce/domain/protocols"
	"github.com/elissonalvesilva/interview-hash/my-ecommerce/tests/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var blackFridayDate = time.Date(2021, 11, 29, 00, 00, 00, 00, time.UTC)

func TestCheckoutController_CheckoutProductsController(t *testing.T) {
	t.Run("Should return a error if body param is invalid json", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)

		mockProductCheckoutRepository := mock.NewMockProductCheckoutRepository(ctrl)
		mockDiscountServiceRepository := mock.NewMockDiscountServiceRepository(ctrl)
		mockValidator := mock.NewMockValidateParam(ctrl)

		productCheckoutUseCase := useCases.NewProductsCheckout(mockProductCheckoutRepository, mockDiscountServiceRepository, blackFridayDate)

		sut := NewCheckoutController(*productCheckoutUseCase, mockValidator)

		productCheckoutRequest := []byte(`{"a1": 1`)

		req, _ := http.NewRequest("POST", "/checkout", bytes.NewBuffer(productCheckoutRequest))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		sut.CheckoutProductsController(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Should return a error if pass incorrect params to body", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)

		mockProductCheckoutRepository := mock.NewMockProductCheckoutRepository(ctrl)
		mockDiscountServiceRepository := mock.NewMockDiscountServiceRepository(ctrl)
		mockValidator := mock.NewMockValidateParam(ctrl)


		productCheckoutUseCase := useCases.NewProductsCheckout(mockProductCheckoutRepository, mockDiscountServiceRepository, blackFridayDate)

		sut := NewCheckoutController(*productCheckoutUseCase, mockValidator)

		mockValidator.EXPECT().ValidateRequestParams(gomock.Any()).Return(errors.New("invalid request params"))

		productCheckoutRequest := []byte(`{"a1": 1}`)

		req, _ := http.NewRequest("POST", "/checkout", bytes.NewBuffer(productCheckoutRequest))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		sut.CheckoutProductsController(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Should return a error if all ids is not in db", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)

		mockProductCheckoutRepository := mock.NewMockProductCheckoutRepository(ctrl)
		mockDiscountServiceRepository := mock.NewMockDiscountServiceRepository(ctrl)
		mockValidator := mock.NewMockValidateParam(ctrl)

		mockProductCheckoutRepository.EXPECT().GetProducts(gomock.Any()).Return([]protocols.ProductToApplyDiscount{})
		mockValidator.EXPECT().ValidateRequestParams(gomock.Any()).Return(nil)

		productCheckoutUseCase := useCases.NewProductsCheckout(mockProductCheckoutRepository, mockDiscountServiceRepository, blackFridayDate)

		sut := NewCheckoutController(*productCheckoutUseCase, mockValidator)

		productCheckoutRequest := []byte(`{"products": [{"id": 17, "quantity": 2}, {"id": 20, "quantity": 2}]}`)

		req, _ := http.NewRequest("POST", "/checkout", bytes.NewBuffer(productCheckoutRequest))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		sut.CheckoutProductsController(w, req)

		var response protocols.CheckoutResponse

		json.Unmarshal(w.Body.Bytes(), &response)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("Should return products applied discount with total values", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)

		mockProductCheckoutRepository := mock.NewMockProductCheckoutRepository(ctrl)
		mockDiscountServiceRepository := mock.NewMockDiscountServiceRepository(ctrl)
		mockValidator := mock.NewMockValidateParam(ctrl)


		gomock.InOrder(
			mockDiscountServiceRepository.EXPECT().GetProductDiscount(gomock.Any()).Return(0.15, nil),
			mockDiscountServiceRepository.EXPECT().GetProductDiscount(gomock.Any()).Return(0.15, nil),
		)
		mockProductCheckoutRepository.EXPECT().GetProducts(gomock.Any()).Return(mock.ProductsToApplyDiscountResponse)
		mockValidator.EXPECT().ValidateRequestParams(gomock.Any()).Return(nil)

		productCheckoutUseCase := useCases.NewProductsCheckout(mockProductCheckoutRepository, mockDiscountServiceRepository, blackFridayDate)

		sut := NewCheckoutController(*productCheckoutUseCase, mockValidator)

		productCheckoutRequest := []byte(`{"products": [{"id": 1, "quantity": 2}, {"id": 3, "quantity": 2}]}`)

		req, _ := http.NewRequest("POST", "/checkout", bytes.NewBuffer(productCheckoutRequest))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		sut.CheckoutProductsController(w, req)

		var response protocols.CheckoutResponse

		json.Unmarshal(w.Body.Bytes(), &response)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, mock.CheckoutResponse, response)
	})
}
