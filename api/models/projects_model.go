package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name        string `gorm:"unique;column:name"`
	Description string `gorm:"column:description"`
	Tasks       []Task `gorm:"foreignKey:ProjectID"`
	UserID      uint   `gorm:"column:user_id"`
	User        User   `gorm:"foreignKey:UserID"`
}
