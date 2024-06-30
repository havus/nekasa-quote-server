package dto

import "github.com/havus/nekasa-quote-server/internal/domain/user"

type UserResponse struct {
	ID                int    `json:"id"`
	FirstName         string `json:"first_name"`
	MiddleName        string `json:"middle_name"`
	LastName          string `json:"last_name"`
	Email             string `json:"email"`
	EncryptedPassword string `json:"-"`
}

type SignUpResponse struct {
	Message string        `json:"message"`
	User    *UserResponse `json:"user"`
}

type SignInResponse struct {
	Message string        `json:"message"`
	User    *UserResponse `json:"user"`
}

type ErrorResponse struct {
	Errors []string `json:"errors"`
}

func NewUserResponse(u *user.User) *UserResponse {
	return &UserResponse{
		ID:                u.ID,
		FirstName:         u.FirstName,
		MiddleName:        u.MiddleName,
		LastName:          u.LastName,
		Email:             u.Email,
		EncryptedPassword: u.EncryptedPassword,
	}
}
