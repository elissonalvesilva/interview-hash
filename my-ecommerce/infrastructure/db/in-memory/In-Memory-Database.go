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
)

var (
	ErrNotFoundItemInDB = errors.New(NotFoundItemInDB)
	ErrNotFoundItemInDBByCondition = errors.New(NotFoundItemInDBByCondition)
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

	return entity.Product{}, ErrNotFoundItemInDB
}

func (db *InMemoryDatabase) GetOne(filter protocols.Filter) (entity.Product, error) {

	for _, product := range db.database {
		if GetValueFromStruct(product, filter.Condition) == filter.ValueToFilter {
			return product, nil
		}
	}

	return entity.Product{}, ErrNotFoundItemInDBByCondition
}

func GetValueFromStruct(product entity.Product, field string) interface{} {
	element := reflect.ValueOf(product).FieldByName(field)

	switch element.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return int(element.Int())
		case reflect.String:
			return element.String()
		case reflect.Float64:
			return element.Float()
		case reflect.Bool:
			return element.Bool()
	}
	return nil
}

