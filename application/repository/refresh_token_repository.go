package repository

import (
	"errors"
	"gin/application/repository/contracts"
	"gin/domain/entities"

	"gorm.io/gorm"
)

type RefreshTokenRepository struct {
	*Repository[entities.RefreshToken]
}

func NewRefreshTokenRepository(db *gorm.DB) contracts.IRefreshTokenRepository {
	return &RefreshTokenRepository{
		Repository: NewRepository[entities.RefreshToken](db),
	}
}

func (r *RefreshTokenRepository) GetByUserID(userID uint) (*entities.RefreshToken, error) {

	var refreshToken entities.RefreshToken

	result := r.db.
		Where("user_id = ?", userID).
		First(&refreshToken)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, result.Error
	}

	return &refreshToken, nil
}
