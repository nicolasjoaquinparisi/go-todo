package models

import (
	"gorm.io/gorm"
)

type TaskStatus string

const (
	ToDo       TaskStatus = "to-do"
	InProgress TaskStatus = "in-progress"
	Complete   TaskStatus = "complete"
)

type Task struct {
	gorm.Model
	Name        string     `gorm:"column:name"`
	Status      TaskStatus `gorm:"type:enum_task_status;default:'ToDo';column:status"`
	Description string     `gorm:"column:description"`
	ProjectID   uint       `gorm:"column:project_id"`
	Project     Project    `gorm:"foreignKey:ProjectID"`
}
