package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-todo/api/routes/auth_routes"
	"go-todo/api/routes/projects_routes"
	"go-todo/api/routes/tasks_routes"
	"go-todo/api/routes/users_routes"
	"go-todo/config"
	"go-todo/database"
	"os"
)

func initialize() {
	config.LoadDotEnv()

	dsn := os.Getenv("DATABASE_URL")
	database.ConnectToDB(dsn)
	database.Migrate()
}

func main() {
	initialize()

	port := os.Getenv("PORT")
	address := fmt.Sprintf("localhost:%v", port)

	router := gin.Default()
	auth_routes.SetupAuthRoutes(router)
	users_routes.SetupUsersRoutes(router)
	projects_routes.SetupProjectsRoutes(router)
	tasks_routes.SetupTasksRoutes(router)

	if err := router.Run(address); err != nil {
		return
	}
}
