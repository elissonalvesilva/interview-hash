package db

import (
	"errors"
	"github.com/elissonalvesilva/interview-hash/my-ecommerce/domain/entity"
)

type InMemoryDatabase struct {
	database []entity.Product
}

type InMemory interface {
	GetByID(id int) (entity.Product, error)
}

const (
	NotFoundItemInDB = "Not found in database"
)

var (
	ErrNotFoundItemInDB = errors.New(NotFoundItemInDB)
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
