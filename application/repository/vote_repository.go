package repository

import (
	"errors"
	"gin/application/repository/contracts"
	"gin/domain/entities"

	"gorm.io/gorm"
)

type VoteRepository struct {
	*Repository[entities.Vote]
}

func NewVoteRepository(db *gorm.DB) contracts.IVoteRepository {
	return &VoteRepository{
		Repository: NewRepository[entities.Vote](db),
	}
}

func (r *VoteRepository) HasAlreadyVoted(PollID uint, UserID uint) (bool, error) {
	var existingVote entities.Vote

	err := r.db.Model(&entities.Vote{}).
		Joins("JOIN poll_categories ON poll_categories.id = votes.poll_category_id").
		Where("poll_categories.poll_id = ? AND votes.user_id = ?", PollID, UserID).
		First(&existingVote).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
