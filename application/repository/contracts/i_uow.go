package contracts

import "gorm.io/gorm"

type IUnitOfWork interface {
	// Begin starts a new transaction
	Begin() (IUnitOfWork, error)

	// Commit commits the transaction
	Commit() error

	// Rollback rolls back the transaction
	Rollback() error

	// DB returns the underlying GORM database
	DB() *gorm.DB

	// Repositories
	// IUserRepository returns the user repository
	IUserRepository() IUserRepository

	// IRefreshTokenRepository returns the refresh token repository
	IRefreshTokenRepository() IRefreshTokenRepository

	// IVoteRepository returns the vote repository
	IVoteRepository() IVoteRepository

	// IPollRepository returns the poll repository
	IPollRepository() IPollRepository

	// IPollCategoryRepository returns the poll category repository
	IPollCategoryRepository() IPollCategoryRepository
}
