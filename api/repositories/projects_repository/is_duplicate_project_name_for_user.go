package projects_repository

import (
	"fmt"
	"github.com/google/uuid"
	"go-todo/api/models"
	"go-todo/database"
)

func IsDuplicateProjectNameForUser(projectName string, userId uuid.UUID) (bool, error) {
	fmt.Println("Invoked: Projects repository IsDuplicateProjectNameForUser")

	var project *models.Project

	result := database.Instance.First(&project, "name = ? AND user_id = ?", projectName, userId)

	if result.Error != nil {
		return false, result.Error
	}

	return project != nil, nil
}
