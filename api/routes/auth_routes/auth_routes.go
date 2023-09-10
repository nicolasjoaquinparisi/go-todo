package auth_routes

import (
	"github.com/gin-gonic/gin"
	"go-todo/api/handlers/auth_handlers"
)

func SetupAuthRoutes(router *gin.Engine) {
	authGroup := router.Group("/api/auth")
	{
		authGroup.POST("/signup", auth_handlers.SignUp)
		authGroup.POST("/signin", auth_handlers.SignIn)
	}
}
