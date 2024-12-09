package contracts

import "gorm.io/gorm"

type IUnitOfWork interface {
	Begin() (IUnitOfWork, error)
	Commit() error
	Rollback() error
	DB() *gorm.DB

	// Repositories
	Users() IUserRepository
	RefreshTokens() IRefreshTokenRepository
}
