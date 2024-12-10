package repository

import (
	"gin/application/repository/contracts"

	"gorm.io/gorm"
)

type UnitOfWork struct {
	db *gorm.DB

	UserRepository         contracts.IUserRepository
	RefreshTokenRepository contracts.IRefreshTokenRepository
}

func NewUnitOfWork(db *gorm.DB) contracts.IUnitOfWork {
	return &UnitOfWork{db: db}
}

func (u *UnitOfWork) Begin() (contracts.IUnitOfWork, error) {
	tx := u.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &UnitOfWork{db: tx}, nil
}

func (u *UnitOfWork) Commit() error {
	return u.db.Commit().Error
}

func (u *UnitOfWork) Rollback() error {
	return u.db.Rollback().Error
}

func (u *UnitOfWork) DB() *gorm.DB {
	return u.db
}

func (u *UnitOfWork) Users() contracts.IUserRepository {
	if u.UserRepository == nil {
		u.UserRepository = NewUserRepository(u.db)
	}
	return u.UserRepository
}

func (u *UnitOfWork) RefreshTokens() contracts.IRefreshTokenRepository {
	if u.RefreshTokenRepository == nil {
		u.RefreshTokenRepository = NewRefreshTokenRepository(u.db)
	}
	return u.RefreshTokenRepository
}
