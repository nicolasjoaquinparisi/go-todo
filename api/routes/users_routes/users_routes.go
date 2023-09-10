package users_routes

import (
	"github.com/gin-gonic/gin"
	"go-todo/api/handlers/users_handlers"
	"go-todo/api/middlewares"
)

func SetupUsersRoutes(router *gin.Engine) {
	usersGroup := router.Group("/api/users")
	{
		usersGroup.GET("/me", middlewares.RequireAuthentication, users_handlers.Me)
	}
}
