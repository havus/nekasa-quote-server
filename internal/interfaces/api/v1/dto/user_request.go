package dto

type SignUpRequest struct {
	Username          string `json:"username" validate:"required"`
	FirstName         string `json:"first_name" validate:"required"`
	MiddleName        string `json:"middle_name"`
	LastName          string `json:"last_name" validate:"required"`
	Email             string `json:"email" validate:"required,email"`
	EncryptedPassword string `json:"password" validate:"required"`
	Phone             string `json:"phone"`
}

type SignInRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
