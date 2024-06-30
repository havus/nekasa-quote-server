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
	validator   *CustomValidator
}

func NewUserHandler(userService user.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
		validator:   NewValidator(),
	}
}

func (h *UserHandler) SignUp(c *gin.Context) {
	var req dto.SignUpRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Errors: []string{err.Error()}})
		return
	}

	// Validation after binding
	errorMessages := h.validator.ValidateStruct(req)
	if errorMessages != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Errors: errorMessages})
		return
	}

	ctx := context.WithValue(context.Background(), "clientIP", c.ClientIP())
	newUser := &user.User{
		Username:          req.Username,
		FirstName:         req.FirstName,
		MiddleName:        req.MiddleName,
		LastName:          req.LastName,
		Email:             req.Email,
		EncryptedPassword: req.EncryptedPassword,
		Phone:             req.Phone,
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
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Errors: []string{err.Error()}})
		return
	}

	// Validation after binding
	errorMessages := h.validator.ValidateStruct(req)
	if errorMessages != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Errors: errorMessages})
		return
	}

	ctx := context.WithValue(context.Background(), "clientIP", c.ClientIP())
	user, err := h.userService.SignIn(ctx, req.Email, req.Password)

	if err != nil {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Errors: []string{"Invalid credentials"}})
		return
	}

	c.JSON(http.StatusOK, dto.SignInResponse{Message: "Sign in successful", User: dto.NewUserResponse(user)})
}
