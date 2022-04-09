package product

import (
	"time"

	"marketplace/role"

	uuid "github.com/satori/go.uuid"
)

type Products struct {
	ID         uuid.UUID
	Name       string
	Detail     string
	Price      int
	Stock      int
	MerchantID uuid.UUID
	Merchant   role.Merchant
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}
