package projects_routes

import (
	"github.com/gin-gonic/gin"
	"go-todo/api/handlers/projects_handlers"
	"go-todo/api/middlewares"
)

func SetupProjectsRoutes(router *gin.Engine) {
	projectsGroup := router.Group("/api/projects")
	{
		projectsGroup.POST("/", middlewares.RequireAuthentication, projects_handlers.Create)
		projectsGroup.GET("/", middlewares.RequireAuthentication, projects_handlers.FindAllByUserId)
		projectsGroup.GET("/:id", middlewares.RequireAuthentication, projects_handlers.FindById)
		projectsGroup.PUT("/:id", middlewares.RequireAuthentication, projects_handlers.Update)
		projectsGroup.DELETE("/:id", middlewares.RequireAuthentication, projects_handlers.HardDelete)
	}
}
