package database

import (
	"fmt"
	"go-todo/api/models"
)

func Migrate() {
	Instance.Exec(`CREATE TYPE enum_task_status AS ENUM ('ToDo', 'InProgress', 'Complete');`)

	err := Instance.AutoMigrate(&models.User{}, &models.Project{}, &models.Task{})

	if err != nil {
		fmt.Print("Failed to generate migration")

		return
	}
}
