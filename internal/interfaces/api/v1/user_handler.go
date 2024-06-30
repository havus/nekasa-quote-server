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
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Errors: []string{"Invalid request"}})
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
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Errors: []string{err.Error()}})
		return
	}

	c.JSON(http.StatusOK, dto.SignUpResponse{Message: "Sign up successful", User: dto.NewUserResponse(newUser)})
}

func (h *UserHandler) SignIn(c *gin.Context) {
	var req dto.SignInRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Errors: []string{"Invalid request"}})
		return
	}

	ctx := context.Background()
	user, err := h.userService.SignIn(ctx, req.Email, req.Password)

	if err != nil {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Errors: []string{"Invalid credentials"}})
		return
	}

	c.JSON(http.StatusOK, dto.SignInResponse{Message: "Sign in successful", User: dto.NewUserResponse(user)})
}
