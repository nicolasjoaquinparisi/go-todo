package projects_routes

import (
	"github.com/gin-gonic/gin"
	"go-todo/api/handlers/projects_handlers"
	"go-todo/api/middlewares"
)

func SetupProjectsRoutes(router *gin.Engine) {
	authGroup := router.Group("/api/projects")
	{
		authGroup.POST("/", middlewares.RequireAuthentication, projects_handlers.Create)
		authGroup.GET("/", middlewares.RequireAuthentication, projects_handlers.FindAllByUserId)
		authGroup.GET("/:id", middlewares.RequireAuthentication, projects_handlers.FindById)
		authGroup.PUT("/:id", middlewares.RequireAuthentication, projects_handlers.Update)
		authGroup.DELETE("/:id", middlewares.RequireAuthentication, projects_handlers.HardDelete)
	}
}
