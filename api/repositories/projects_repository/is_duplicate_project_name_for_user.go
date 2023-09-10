package projects_repository

import (
	"go-todo/api/models"
	"go-todo/database"
)

func IsDuplicateProjectNameForUser(projectName string, userId uint) (bool, error) {
	var project *models.Project

	result := database.DB.First(&project, "name = ? AND user_id = ?", projectName, userId)

	// return query error
	if result.Error != nil {
		return false, result.Error
	}

	return project != nil, nil
}
