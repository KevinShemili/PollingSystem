package contracts

import (
	"gin/domain/entities"
	"time"
)

type IPollRepository interface {
	IRepository[entities.Poll]

	GetPollWithVotes(pollID uint) (*entities.Poll, error)
	GetExpiredPolls(currentTime time.Time) ([]*entities.Poll, error)
	GetPollWithCategories(pollID uint) (*entities.Poll, error)
}
