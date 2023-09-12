package projects_responses

import "go-todo/api/models"

type FindByIdResponse struct {
	ID          uint          `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Tasks       []models.Task `json:"tasks"`
}
