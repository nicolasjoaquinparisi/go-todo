package routes

import (
	"github.com/gin-gonic/gin"
	"go-todo/api/handlers"
	"go-todo/api/middlewares"
)

func SetupUsersRoutes(router *gin.Engine) {
	usersGroup := router.Group("/api/users")
	{
		usersGroup.GET("/me", middlewares.RequireAuthentication, handlers.Me)
	}
}
