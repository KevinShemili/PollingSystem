package entities

import "gorm.io/gorm"

type PollCategory struct {
	gorm.Model

	Name string

	// fk
	PollID uint

	// relations
	Poll  Poll   `gorm:"foreignKey:PollID;constraint:OnDelete:SET NULL;"`
	Votes []Vote `gorm:"foreignKey:PollCategoryID;constraint:OnDelete:SET NULL;"`
}
