package repository

import (
	"gin/application/repository/contracts"

	"gorm.io/gorm"
)

type UnitOfWork struct {
	Database *gorm.DB

	UserRepository         contracts.IUserRepository
	RefreshTokenRepository contracts.IRefreshTokenRepository
	VoteRepository         contracts.IVoteRepository
	PollRepository         contracts.IPollRepository
	PollCategoryRepository contracts.IPollCategoryRepository
}

func NewUnitOfWork(database *gorm.DB) contracts.IUnitOfWork {
	return &UnitOfWork{Database: database}
}

func (unitOfWork *UnitOfWork) Begin() (contracts.IUnitOfWork, error) {
	tx := unitOfWork.Database.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &UnitOfWork{Database: tx}, nil
}

func (unitOfWork *UnitOfWork) Commit() error {
	return unitOfWork.Database.Commit().Error
}

func (unitOfWork *UnitOfWork) Rollback() error {
	return unitOfWork.Database.Rollback().Error
}

func (unitOfWork *UnitOfWork) DB() *gorm.DB {
	return unitOfWork.Database
}

func (unitOfWork *UnitOfWork) IUserRepository() contracts.IUserRepository {
	if unitOfWork.UserRepository == nil {
		unitOfWork.UserRepository = NewUserRepository(unitOfWork.Database)
	}
	return unitOfWork.UserRepository
}

func (unitOfWork *UnitOfWork) IRefreshTokenRepository() contracts.IRefreshTokenRepository {
	if unitOfWork.RefreshTokenRepository == nil {
		unitOfWork.RefreshTokenRepository = NewRefreshTokenRepository(unitOfWork.Database)
	}
	return unitOfWork.RefreshTokenRepository
}

func (unitOfWork *UnitOfWork) IVoteRepository() contracts.IVoteRepository {
	if unitOfWork.VoteRepository == nil {
		unitOfWork.VoteRepository = NewVoteRepository(unitOfWork.Database)
	}
	return unitOfWork.VoteRepository
}

func (unitOfWork *UnitOfWork) IPollRepository() contracts.IPollRepository {
	if unitOfWork.PollRepository == nil {
		unitOfWork.PollRepository = NewPollRepository(unitOfWork.Database)
	}
	return unitOfWork.PollRepository
}

func (unitOfWork *UnitOfWork) IPollCategoryRepository() contracts.IPollCategoryRepository {
	if unitOfWork.PollCategoryRepository == nil {
		unitOfWork.PollCategoryRepository = NewPollCategoryRepository(unitOfWork.Database)
	}
	return unitOfWork.PollCategoryRepository
}
