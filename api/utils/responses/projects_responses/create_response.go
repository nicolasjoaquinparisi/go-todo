package projects_responses

import (
	"github.com/google/uuid"
	"go-todo/api/models"
)

type CreateHandlerResponse struct {
	ID          uuid.UUID     `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Tasks       []models.Task `json:"tasks"`
	UserID      uuid.UUID     `json:"user_id"`
}
