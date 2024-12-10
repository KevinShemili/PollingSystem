package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model

	FirstName    string
	LastName     string
	Email        string `gorm:"unique"`
	Age          int
	PasswordHash string
	IsDeleted    bool

	RefreshTokens []RefreshToken `gorm:"foreignKey:UserID"`
}
