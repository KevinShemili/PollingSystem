package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName        string
	LastName         string
	Email            string `gorm:"unique"`
	IsEmailConfirmed bool
	PasswordHash     string
	Age              int
	IsDeleted        bool
}
