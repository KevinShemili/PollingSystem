package contracts

import "gin/domain/entities"

type IVoteRepository interface {
	IRepository[entities.Vote]
}
