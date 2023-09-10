package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email     string    `gorm:"unique;column:email"`
	FirstName string    `gorm:"column:first_name"`
	LastName  string    `gorm:"column:last_name"`
	Password  string    `gorm:"column:password"`
	Projects  []Project `gorm:"foreignKey:UserID"`
}
