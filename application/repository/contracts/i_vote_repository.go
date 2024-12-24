package contracts

import "gin/domain/entities"

type IVoteRepository interface {
	IRepository[entities.Vote]

	HasAlreadyVoted(PollID uint, UserID uint) (bool, error)
}
