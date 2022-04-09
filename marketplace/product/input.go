package product

import (
	"marketplace/role"

	uuid "github.com/satori/go.uuid"
)

type CreateProductInput struct {
	ID       uuid.UUID
	Name     string `json:"name"`
	Detail   string `json:"details"`
	Price    int    `json:"price"`
	Stock    int    `json:"stock"`
	Merchant role.Merchant
}
