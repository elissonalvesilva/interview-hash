package protocols

import "github.com/elissonalvesilva/interview-hash/my-ecommerce/domain/entity"

type Filter struct {
	Condition string
	ValueToFilter interface{}
}

type InMemory interface {
	GetByID(id int) (entity.Product, error)
	GetOne(filter Filter) (entity.Product, error)
}
