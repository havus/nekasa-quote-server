package v1

import (
	"context"
	"net/http"

	"github.com/havus/nekasa-quote-server/internal/domain/user"
	"github.com/havus/nekasa-quote-server/internal/interfaces/api/v1/dto"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService user.UserService
}

func NewUserHandler(userService user.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) SignUp(c *gin.Context) {
	var req dto.SignUpRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	ctx := context.Background()
	newUser := &user.User{
		FirstName:         req.FirstName,
		MiddleName:        req.MiddleName,
		LastName:          req.LastName,
		Email:             req.Email,
		EncryptedPassword: req.EncryptedPassword,
	}

	if err := h.userService.SignUp(ctx, newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Sign up successful", "user": dto.NewUserResponse(newUser)})
}

func (h *UserHandler) SignIn(c *gin.Context) {
	var req dto.SignInRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	ctx := context.Background()
	user, err := h.userService.SignIn(ctx, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Sign in successful", "user": dto.NewUserResponse(user)})
}
