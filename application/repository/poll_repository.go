package repository

import (
	"errors"
	"gin/application/repository/contracts"
	"gin/application/utility"
	"gin/domain/entities"
	"time"

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
	err := r.db.
		Preload("Categories.Votes").
		First(&poll, pollID).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &poll, nil
}

func (r *PollRepository) GetPollWithCategories(pollID uint) (*entities.Poll, error) {

	var poll entities.Poll
	err := r.db.
		Preload("Categories").
		First(&poll, pollID).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &poll, nil
}

func (r *PollRepository) GetExpiredPolls(currentTime time.Time) ([]*entities.Poll, error) {
	var polls []*entities.Poll

	err := r.db.
		Where("is_ended = ? AND expires_at < ?", false, currentTime).
		Preload("Creator").
		Preload("Categories.Votes").
		Find(&polls).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return polls, nil
}

func (r *PollRepository) GetPollsPaginated(parameters utility.QueryParams, showActiveOnly bool) (utility.PaginatedResponse[entities.Poll], error) {

	db := r.db.
		Model(&entities.Poll{}).
		Preload("Categories.Votes")

	if showActiveOnly {
		db = db.Where("is_ended = ?", false)
	}

	return utility.PaginateAndFilter[entities.Poll](db, parameters)
}

func (r *PollRepository) GetPollsByUserPaginated(userID uint, parameters utility.QueryParams, showActiveOnly bool) (utility.PaginatedResponse[entities.Poll], error) {

	db := r.db.
		Model(&entities.Poll{}).
		Preload("Categories.Votes")

	if showActiveOnly {
		db = db.Where("is_ended = ?", false)
	}

	return utility.PaginateAndFilter[entities.Poll](db, parameters)
}
