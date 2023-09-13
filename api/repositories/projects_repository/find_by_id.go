package projects_repository

import (
	"github.com/google/uuid"
	"go-todo/api/models"
	"go-todo/database"
	"gorm.io/gorm"
)

func FindById(id uuid.UUID) (*models.Project, error) {
	var project *models.Project

	result := database.Instance.Model(&models.Project{}).
		Preload("Tasks", func(db *gorm.DB) *gorm.DB {
			return db.Select("name")
		}).
		Find(&project, "id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}

	return project, nil
}
