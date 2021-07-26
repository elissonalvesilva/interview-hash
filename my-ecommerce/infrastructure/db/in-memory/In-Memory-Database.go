package in_memory

import (
	"errors"
	"github.com/elissonalvesilva/interview-hash/my-ecommerce/domain/entity"
	"github.com/elissonalvesilva/interview-hash/my-ecommerce/infrastructure/db/in-memory/protocols"
	"reflect"
)

type InMemoryDatabase struct {
	database []entity.Product
}

const (
	NotFoundItemInDB = "Not found in database"
	NotFoundItemInDBByCondition = "Not found in database with this condition"
	NotFoundConditionInStruct = "Not found condition in struct"
)

var (
	ErrNotFoundItemInDB = errors.New(NotFoundItemInDB)
	ErrNotFoundItemInDBByCondition = errors.New(NotFoundItemInDBByCondition)
	ErrNotFoundConditionInStruct = errors.New(NotFoundConditionInStruct)
)

func NewInMemoryDatabase(db []entity.Product) *InMemoryDatabase {
	return &InMemoryDatabase{
		database: db,
	}
}

func (db *InMemoryDatabase) GetByID(id int) (entity.Product, error) {
	for _, product := range db.database {
		if id == int(product.ID) {
			return product, nil
		}
	}

	return entity.Product{}, ErrNotFoundItemInDBByCondition
}

func (db *InMemoryDatabase) GetOne(filter protocols.Filter) (entity.Product, error) {
	return entity.Product{}, ErrNotFoundItemInDB
}

func GetNameStruct(elementToGet string) (string, error) {
	element := reflect.ValueOf(&entity.Product{}).Elem()

	for i := 0; i < element.NumField(); i++ {
		name := element.Type().Field(i).Name

		if name == elementToGet {
			return name, nil
		}
	}

	return "", ErrNotFoundConditionInStruct
}
