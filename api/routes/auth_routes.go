package routes

import (
	"github.com/gin-gonic/gin"
	"go-todo/api/handlers"
)

func SetupAuthRoutes(router *gin.Engine) {
	authGroup := router.Group("/api/auth")
	{
		authGroup.POST("/signup", handlers.SignUp)
		authGroup.POST("/signin", handlers.SignIn)
	}
}
