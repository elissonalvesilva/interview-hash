package in_memory

import (
	"github.com/elissonalvesilva/interview-hash/my-ecommerce/domain/entity"
	"github.com/elissonalvesilva/interview-hash/my-ecommerce/tests/mock"
	"github.com/stretchr/testify/assert"
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

func TestGetNameStruct(t *testing.T) {
	t.Run("Should return a error if condition is not found in struct", func(t *testing.T) {
		t.Parallel()

		expectedResponse := ErrNotFoundConditionInStruct

		response, errorResponse := GetNameStruct("x")
		assert.Equal(t, "", response)
		assert.NotNil(t, errorResponse)
		assert.Errorf(t, expectedResponse, response)
	})

	t.Run("Should return a string with name of Struct key if condition exists", func(t *testing.T) {
		t.Parallel()

		expectedResponse := "ID"

		response, errorResponse := GetNameStruct("ID")
		assert.Nil(t, errorResponse)
		assert.Equal(t, expectedResponse, response)

	})
}
