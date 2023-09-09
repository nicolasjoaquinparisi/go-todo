package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-todo/api/routes"
	"go-todo/config"
	"go-todo/database"
	"os"
)

func initialize() {
	config.LoadDotEnv()

	database.ConnectToDB()
	database.Migrate()
}

func main() {
	initialize()

	port := os.Getenv("PORT")
	address := fmt.Sprintf("localhost:%v", port)

	router := gin.Default()
	routes.SetupAuthRoutes(router)

	if err := router.Run(address); err != nil {
		return
	}
}
