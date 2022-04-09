package user

import uuid "github.com/satori/go.uuid"

type UserFormatter struct {
	ID     uuid.UUID `json:"id"`
	RoleId int       `json:"role_id"`
	Name   string    `json:"name"`
	Email  string    `json:"email"`
	Token  string    `json:"token"`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:     user.ID,
		RoleId: user.RoleId,
		Name:   user.FirstName + " " + user.LastName,
		Email:  user.Email,
		Token:  token,
	}

	return formatter
}
