package role

import (
	"marketplace/user"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Merchant struct {
	ID           uuid.UUID
	MerchantName string
	Address      string
	UserId       uuid.UUID
	User         user.User
	RoleId       int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}

type Customer struct {
	ID        uuid.UUID
	Address   string
	UserId    uuid.UUID
	RoleId    int
	User      user.User
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
