package main

import (
	"bytes"
	"encoding/json"
	"github.com/elissonalvesilva/interview-hash/my-ecommerce/domain/protocols"
	inMemoryDB "github.com/elissonalvesilva/interview-hash/my-ecommerce/infrastructure/db/in-memory"
	errorToResponse "github.com/elissonalvesilva/interview-hash/my-ecommerce/presenters/error"
	presentersProtocol "github.com/elissonalvesilva/interview-hash/my-ecommerce/presenters/protocols"
	"github.com/elissonalvesilva/interview-hash/my-ecommerce/tests/mock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func setup(t *testing.T, needsDiscountService bool) (*httptest.ResponseRecorder, *mux.Router) {
	db := mock.Products
	var blackFridayDate time.Time

	database := inMemoryDB.NewInMemoryDatabase(db)

	blackFridayDate = time.Date(time.Now().Year(), 11, 29, 00, 00, 00, 00, time.UTC)

	router := mux.NewRouter()

	controller := mock.SetupController(t, database, blackFridayDate, needsDiscountService)

	router.HandleFunc("/checkout", controller.CheckoutProductsController).Methods("POST")
	router.HandleFunc("/health", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode("OK")
		return
	}).Methods("GET")
	//The response recorder used to record HTTP responses
	writer := httptest.NewRecorder()
	return writer, router
}

func TestE2ECheckoutProduct(t *testing.T) {
	t.Run("Should return ok if application is health", func(t *testing.T) {
		t.Parallel()
		needsDiscountService := false
		writer, router := setup(t, needsDiscountService)

		req, _ := http.NewRequest("GET", "/health", nil)

		router.ServeHTTP(writer, req)
		var response string

		json.Unmarshal(writer.Body.Bytes(), &response)

		assert.Equal(t, http.StatusOK, writer.Code)
		assert.Equal(t, "OK", response)
	})

	t.Run("Should return response with total values and products on success", func(t *testing.T) {
		t.Parallel()
		needsDiscountService := true

		writer, router := setup(t, needsDiscountService)
		productCheckoutRequest := []byte(`{"products": [{"id": 1, "quantity": 2}, {"id": 3, "quantity": 2}]}`)

		req, _ := http.NewRequest("POST", "/checkout", bytes.NewBuffer(productCheckoutRequest))
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(writer, req)
		var response protocols.CheckoutResponse

		json.Unmarshal(writer.Body.Bytes(), &response)

		assert.Equal(t, http.StatusOK, writer.Code)
		assert.Equal(t, mock.CheckoutResponse, response)
	})

	t.Run("Should return error if json is invalid", func(t *testing.T) {
		t.Parallel()
		needsDiscountService := false
		writer, router := setup(t, needsDiscountService)
		productCheckoutRequest := []byte(`{"invalid"`)

		req, _ := http.NewRequest("POST", "/checkout", bytes.NewBuffer(productCheckoutRequest))
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(writer, req)
		var response presentersProtocol.ErrorResponse

		json.Unmarshal(writer.Body.Bytes(), &response)
		assert.Equal(t, http.StatusBadRequest, writer.Code)
		assert.Equal(t, errorToResponse.InvalidJSONParamMessage ,response.Message)
	})

	t.Run("Should return error if is invalid body param", func(t *testing.T) {
		t.Parallel()
		needsDiscountService := false
		writer, router := setup(t, needsDiscountService)
		productCheckoutRequest := []byte(`{"x": 1}`)

		req, _ := http.NewRequest("POST", "/checkout", bytes.NewBuffer(productCheckoutRequest))
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(writer, req)
		var response presentersProtocol.ErrorResponse

		json.Unmarshal(writer.Body.Bytes(), &response)

		assert.Equal(t, http.StatusBadRequest, writer.Code)
		assert.NotNil(t, response.Message)
		assert.NotNil(t, response.Stack)

	})

	t.Run("Should return not found if not found all products id in db", func(t *testing.T) {
		t.Parallel()
		needsDiscountService := false
		writer, router := setup(t, needsDiscountService)
		productCheckoutRequest := []byte(`{"products": [{"id": 17, "quantity": 2}, {"id": 19, "quantity": 2}]}`)

		req, _ := http.NewRequest("POST", "/checkout", bytes.NewBuffer(productCheckoutRequest))
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(writer, req)
		var response presentersProtocol.ErrorResponse

		json.Unmarshal(writer.Body.Bytes(), &response)

		assert.Equal(t, http.StatusNotFound, writer.Code)
		assert.NotNil(t, response.Message)
		assert.Equal(t, "Not found ids: 17,19", response.Message)
	})
}