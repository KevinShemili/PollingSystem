package entities

import (
	"time"

	"gorm.io/gorm"
)

type RefreshToken struct {
	gorm.Model

	Token     string
	Expiry    time.Time
	JWTToken  string
	IsDeleted bool `gorm:"default:false"`

	// fk
	UserID uint

	// relations
	User User `gorm:"foreignKey:UserID;constraint:OnDelete:SET NULL;"`
}
