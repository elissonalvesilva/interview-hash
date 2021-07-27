package in_memory

import (
	"github.com/elissonalvesilva/interview-hash/my-ecommerce/domain/entity"
	"github.com/elissonalvesilva/interview-hash/my-ecommerce/infrastructure/db/in-memory/protocols"
	"github.com/elissonalvesilva/interview-hash/my-ecommerce/tests/mock"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestInMemoryDatabase_GetByID(t *testing.T) {
	t.Run("Should return a error if not found item in database", func(t *testing.T) {
		expectedResponse := ErrNotFoundItemInDB
		database := mock.Products

		sut := NewInMemoryDatabase(database)
		response, errorResponse := sut.GetByID(5)

		assert.Equal(t, entity.Product{}, response)
		assert.NotNil(t, errorResponse)
		assert.Equal(t, expectedResponse, errorResponse)
	})

	t.Run("Should return a product in database", func(t *testing.T) {
		expectedResponse := mock.Product1
		database := mock.Products

		sut := NewInMemoryDatabase(database)
		response, errorResponse := sut.GetByID(1)

		assert.Nil(t, errorResponse)
		assert.Equal(t, expectedResponse, response)
	})
}

func TestInMemoryDatabase_GetOne(t *testing.T) {
	t.Run("Should return error if item not found in database using filter", func(t *testing.T) {
		expectedResponse := ErrNotFoundItemInDBByCondition
		database := mock.Products

		sut := NewInMemoryDatabase(database)
		response, errorResponse := sut.GetOne(protocols.Filter{Condition: "ID", ValueToFilter: 5})

		assert.Equal(t, entity.Product{}, response)
		assert.NotNil(t, errorResponse)
		assert.Equal(t, expectedResponse, errorResponse)
	})

	t.Run("Should return product by filter", func(t *testing.T) {
		database := mock.Products

		sut := NewInMemoryDatabase(database)
		response, errorResponse := sut.GetOne(protocols.Filter{Condition: "ID", ValueToFilter: 1})

		assert.Equal(t, mock.Product1, response)
		assert.Nil(t, errorResponse)
	})

	t.Run("Should return product by filter equals to gift true", func(t *testing.T) {
		database := mock.Products

		sut := NewInMemoryDatabase(database)
		response, errorResponse := sut.GetOne(protocols.Filter{Condition: "IsGift", ValueToFilter: true})

		assert.Equal(t, mock.Product3, response)
		assert.Nil(t, errorResponse)
	})
}

func TestGetValueFromStruct(t *testing.T) {
	t.Run("Should return element equals to int", func(t *testing.T) {
		response := GetValueFromStruct(mock.Product1, "ID")
		expectedInt64 := reflect.TypeOf(mock.Product1.ID).Kind()
		assert.IsType(t, expectedInt64, reflect.TypeOf(response).Kind())
	})

	t.Run("Should return element equals to string", func(t *testing.T) {
		response := GetValueFromStruct(mock.Product1, "Title")
		expectedString := reflect.TypeOf(mock.Product1.Title).Kind()
		assert.IsType(t, expectedString, reflect.TypeOf(response).Kind())
	})

	t.Run("Should return element equals to float", func(t *testing.T) {
		response := GetValueFromStruct(mock.Product1, "Amount")
		expectedFloat64 := reflect.TypeOf(mock.Product1.Amount).Kind()
		assert.IsType(t, expectedFloat64, reflect.TypeOf(response).Kind())
	})

	t.Run("Should return element equals to bool", func(t *testing.T) {
		response := GetValueFromStruct(mock.Product1, "IsGift")
		expectedBool := reflect.TypeOf(mock.Product1.IsGift).Kind()
		assert.IsType(t, expectedBool, reflect.TypeOf(response).Kind())
	})
}
