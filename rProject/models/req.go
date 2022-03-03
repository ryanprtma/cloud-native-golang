package models

import (
	"time"
)

type (
	Req struct {
		ID        int       `json:"id"`
		Email     string    `json:"email"`
		Text      string    `json:"text"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
