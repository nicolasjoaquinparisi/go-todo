package projects_responses

import "github.com/google/uuid"

type FindAllByUserIdResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}
