package mock

import "github.com/elissonalvesilva/interview-hash/my-ecommerce/domain/entity"

var Product1 = entity.Product{
	ID: 1,
	Title: "Ergonomic Wooden Pants",
	Description: "Deleniti beatae porro.",
	Amount: 15157,
	IsGift: false,
}

var Product2 = entity.Product{
	ID: 2,
	Title: "Ergonomic Cotton Keyboard",
	Description: "Iste est ratione excepturi repellendus adipisci qui.",
	Amount: 93811,
	IsGift: false,
}

var Product3 = entity.Product{
	ID: 3,
	Title: "Gorgeous Cotton Chips",
	Description: "Nulla rerum tempore rem.",
	Amount: 60356,
	IsGift: true,
}

var Products = []entity.Product{
	Product1,
	Product2,
	Product3,
}
