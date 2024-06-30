package v1

import (
	"github.com/havus/nekasa-quote-server/internal/domain/user"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, userService user.UserService) {
	v1 := router.Group("/v1")
	userHandler := NewUserHandler(userService)

	v1.POST("/sign-in", userHandler.SignIn)
	v1.POST("/sign-up", userHandler.SignUp)
}
