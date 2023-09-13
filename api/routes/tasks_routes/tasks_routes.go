package tasks_routes

import (
	"github.com/gin-gonic/gin"
	"go-todo/api/handlers/tasks_handlers"
	"go-todo/api/middlewares"
)

func SetupTasksRoutes(router *gin.Engine) {
	tasksGroup := router.Group("/api/tasks")
	{
		tasksGroup.POST("/", middlewares.RequireAuthentication, tasks_handlers.Create)
		//tasksGroup.GET("/", middlewares.RequireAuthentication, tasks_handlers.FindAllByUserId)
		//tasksGroup.GET("/:id", middlewares.RequireAuthentication, tasks_handlers.FindById)
		//tasksGroup.PUT("/:id", middlewares.RequireAuthentication, tasks_handlers.Update)
		//tasksGroup.DELETE("/:id", middlewares.RequireAuthentication, tasks_handlers.HardDelete)
	}
}
