package projects_repository

import (
	"go-todo/api/models"
	"go-todo/database"
)

func FindAllByUserId(userId uint) ([]*models.Project, error) {
	var projects []*models.Project

	result := database.DB.Find(&projects, "user_id = ?", userId)

	if result.Error != nil {
		return nil, result.Error
	}

	return projects, nil
}
