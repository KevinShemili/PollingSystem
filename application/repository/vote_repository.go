package repository

import (
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
