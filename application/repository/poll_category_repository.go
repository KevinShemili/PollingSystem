package repository

import (
	"errors"
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

func (r *PollCategoryRepository) HasCategory(PollID uint, PollCategoryID uint) (bool, error) {

	var category entities.PollCategory
	err := r.db.
		Model(&entities.PollCategory{}).
		Where("poll_id = ? AND id = ?", PollID, PollCategoryID).
		First(&category).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
