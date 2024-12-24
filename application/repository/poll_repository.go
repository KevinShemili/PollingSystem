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

func (r *PollRepository) GetPollWithVotes(pollID uint) (*entities.Poll, error) {

	var poll entities.Poll
	err := r.db.Preload("Categories.Votes").First(&poll, pollID).Error

	if err != nil {
		return nil, err
	}

	return &poll, nil
}
