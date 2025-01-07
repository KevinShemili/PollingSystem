package contracts

import "gin/domain/entities"

type IRefreshTokenRepository interface {
	IRepository[entities.RefreshToken]

	// GetByUserID returns the refresh token for a given user ID
	GetByUserID(userID uint) (*entities.RefreshToken, error)
}
