package entities

import (
	"time"

	"gorm.io/gorm"
)

type Poll struct {
	gorm.Model

	Title       string
	Description string
	ExpiresAt   time.Time
	IsEnded     bool `gorm:"default:false"`

	// fk
	CreatorID uint

	// relations
	Creator    User           `gorm:"foreignKey:CreatorID;constraint:OnDelete:SET NULL;"`
	Categories []PollCategory `gorm:"foreignKey:PollID;constraint:OnDelete:SET NULL;"`
}
