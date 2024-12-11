package repository

import (
	"gin/application/repository/contracts"
	"gin/domain/entities"

	"gorm.io/gorm"
)

type PollRepository struct {
	*Repository[entities.Poll]
}

func NewPollRepository(db *gorm.DB) contracts.IPollRepository {
	return &PollRepository{
		Repository: NewRepository[entities.Poll](db),
	}
}
