package user

type RegisterUserInput struct {
	FirstName    string `json:"firstname" binding:"required"`
	LastName     string `json:"lastname" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
	HashPassword string `json:"hashpassword" binding:"required"`
	RoleId       int    `gorm:"default:0"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type CheckEmailInput struct {
	Email string `json:"email" binding:"required,email"`
}
