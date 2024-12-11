package contracts

import "gorm.io/gorm"

type IUnitOfWork interface {
	Begin() (IUnitOfWork, error)
	Commit() error
	Rollback() error
	DB() *gorm.DB

	// Repositories
	IUserRepository() IUserRepository
	IRefreshTokenRepository() IRefreshTokenRepository
	IVoteRepository() IVoteRepository
	IPollRepository() IPollRepository
	IPollCategoryRepository() IPollCategoryRepository
}
