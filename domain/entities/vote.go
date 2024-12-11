package entities

import "gorm.io/gorm"

type Vote struct {
	gorm.Model

	// fk
	UserID         uint
	PollCategoryID uint

	// relations
	User         User         `gorm:"foreignKey:UserID;constraint:OnDelete:SET NULL;"`
	PollCategory PollCategory `gorm:"foreignKey:PollCategoryID;constraint:OnDelete:SET NULL;"`
}
