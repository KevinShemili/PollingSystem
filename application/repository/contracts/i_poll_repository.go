package contracts

import (
	"gin/application/utility"
	"gin/domain/entities"
	"time"
)

type IPollRepository interface {
	IRepository[entities.Poll]

	GetPollWithVotes(pollID uint) (*entities.Poll, error)
	GetExpiredPolls(currentTime time.Time) ([]*entities.Poll, error)
	GetPollWithCategories(pollID uint) (*entities.Poll, error)
	GetPollsPaginated(parameters utility.QueryParams, showActiveOnly bool) (utility.PaginatedResponse[entities.Poll], error)
	GetPollsByUserPaginated(userID uint, parameters utility.QueryParams, showActiveOnly bool) (utility.PaginatedResponse[entities.Poll], error)
}
