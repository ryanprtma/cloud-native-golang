package user

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID             uuid.UUID
	FirstName      string
	LastName       string
	Email          string
	AvatarFileName string
	HashPassword   string
	RoleId         int
	Token          string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	// DeletedAt      time.Time
}
