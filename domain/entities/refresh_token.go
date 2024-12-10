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
	IsDeleted bool

	UserID uint
	User   User `gorm:"constraint:OnDelete:SET NULL;"`
}
