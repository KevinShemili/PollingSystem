package repository

import (
	"errors"
	"gin/application/repository/contracts"
	"gin/domain/entities"

	"gorm.io/gorm"
)

type UserRepository struct {
	*Repository[entities.User]
}

func NewUserRepository(db *gorm.DB) contracts.IUserRepository {
	return &UserRepository{
		Repository: NewRepository[entities.User](db),
	}
}

func (r *UserRepository) GetByEmail(email string) (*entities.User, error) {

	var user entities.User

	result := r.db.
		Where("email = ?", email).
		First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, result.Error
	}

	return &user, nil
}
