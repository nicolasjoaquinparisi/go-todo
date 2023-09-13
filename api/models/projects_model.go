package models

import "github.com/google/uuid"

type Project struct {
	BaseModel
	Name        string    `gorm:"column:name"`
	Description string    `gorm:"column:description"`
	UserID      uuid.UUID `gorm:"type:uuid;column:user_id;not null;"`
	User        User      `gorm:"foreignKey:UserID"`
	Tasks       []Task    `gorm:"foreignKey:ID"`
}
