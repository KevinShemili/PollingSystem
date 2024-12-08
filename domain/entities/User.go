package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	Email            string `gorm:"unique" json:"email"`
	IsEmailConfirmed bool   `json:"is_email_confirmed"`
	PasswordHash     string `json:"password_hash"`
	PasswordSalt     string `json:"password_salt"`
	Age              int    `json:"age"`
}
