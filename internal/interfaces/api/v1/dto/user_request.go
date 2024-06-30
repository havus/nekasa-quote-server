package dto

type SignUpRequest struct {
	FirstName         string `json:"first_name" binding:"required"`
	MiddleName        string `json:"middle_name"`
	LastName          string `json:"last_name" binding:"required"`
	Email             string `json:"email" binding:"required"`
	EncryptedPassword string `json:"password" binding:"required"`
}

type SignInRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
