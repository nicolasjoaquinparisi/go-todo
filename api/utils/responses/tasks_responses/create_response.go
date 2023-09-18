package tasks_responses

import (
	"github.com/google/uuid"
	"go-todo/api/models"
)

type CreateHandlerResponse struct {
	ID          uuid.UUID         `json:"id"`
	Name        string            `json:"name"`
	Status      models.TaskStatus `json:"status"`
	Description string            `json:"description"`
	ProjectID   uuid.UUID         `json:"project_id"`
}
