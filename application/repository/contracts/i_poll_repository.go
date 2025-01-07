package contracts

import (
	"gin/application/utility"
	"gin/domain/entities"
	"time"
)

type IPollRepository interface {
	IRepository[entities.Poll]

	// GetPollWithVotes returns a poll with preloaded categories and votes
	GetPollWithVotes(pollID uint) (*entities.Poll, error)

	// GetExpiredPolls returns all polls that have expired
	GetExpiredPolls(currentTime time.Time) ([]*entities.Poll, error)

	// GetPollWithCategories returns a poll with preloaded categories
	GetPollWithCategories(pollID uint) (*entities.Poll, error)

	// GetPollsPaginated returns a paginated list of polls with preloaded categories and votes. If showActiveOnly is true, only active polls are returned
	GetPollsPaginated(parameters utility.QueryParams, showActiveOnly bool) (utility.PaginatedResponse[entities.Poll], error)

	// GetPollsByUserPaginated returns a paginated list of polls created by a user with preloaded categories and votes. If showActiveOnly is true, only active polls are returned
	GetPollsByUserPaginated(userID uint, parameters utility.QueryParams, showActiveOnly bool) (utility.PaginatedResponse[entities.Poll], error)
}
