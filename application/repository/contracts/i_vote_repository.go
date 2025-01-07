package contracts

import "gin/domain/entities"

type IVoteRepository interface {
	IRepository[entities.Vote]

	// HasAlreadyVoted checks if a user has already voted in a poll
	HasAlreadyVoted(PollID uint, UserID uint) (bool, error)
}
