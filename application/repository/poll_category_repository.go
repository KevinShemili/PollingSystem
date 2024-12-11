package repository

import (
	"gin/application/repository/contracts"
	"gin/domain/entities"

	"gorm.io/gorm"
)

type PollCategoryRepository struct {
	*Repository[entities.PollCategory]
}

func NewPollCategoryRepository(db *gorm.DB) contracts.IPollCategoryRepository {
	return &PollCategoryRepository{
		Repository: NewRepository[entities.PollCategory](db),
	}
}
