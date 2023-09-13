package projects_responses

import (
	"github.com/google/uuid"
	"go-todo/api/models"
)

type FindByIdResponse struct {
	ID          uuid.UUID     `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Tasks       []models.Task `json:"tasks"`
}
