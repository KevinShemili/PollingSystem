package contracts

import "gin/domain/entities"

type IRefreshTokenRepository interface {
	IRepository[entities.RefreshToken]

	GetByUserID(userID uint) (*entities.RefreshToken, error)
}
