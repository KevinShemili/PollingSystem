package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model

	FirstName    string
	LastName     string
	Email        string `gorm:"unique"`
	Age          int
	PasswordHash string
	IsDeleted    bool `gorm:"default:false"`

	// relations
	CreatedPolls  []Poll         `gorm:"foreignKey:CreatorID;constraint:OnDelete:SET NULL;"`
	RefreshTokens []RefreshToken `gorm:"foreignKey:UserID;constraint:OnDelete:SET NULL;"`
}
