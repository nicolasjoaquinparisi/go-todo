package models

type User struct {
	BaseModel
	Email     string `gorm:"unique;column:email"`
	FirstName string `gorm:"column:first_name"`
	LastName  string `gorm:"column:last_name"`
	Password  string `gorm:"column:password"`
	Projects  []Project
}
