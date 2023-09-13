package models

import "github.com/google/uuid"

type TaskStatus string

const (
	ToDo       TaskStatus = "to-do"
	InProgress TaskStatus = "in-progress"
	Complete   TaskStatus = "complete"
)

type Task struct {
	BaseModel
	Name        string     `gorm:"column:name"`
	Status      TaskStatus `gorm:"type:enum_task_status;default:'ToDo';column:status"`
	Description string     `gorm:"column:description"`
	ProjectID   uuid.UUID  `gorm:"column:project_id"`
	Project     Project    `gorm:"foreignKey:ProjectID"`
}
