package repository

import (
	"gin/domain/entities"

	"gorm.io/gorm"
)

type UserRepository struct {
	*Repository[entities.User]
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		Repository: NewRepository[entities.User](db),
	}
}

func (r *UserRepository) FindByEmail(email string) (*entities.User, error) {

	var user entities.User

	result := r.db.Where("email = ?", email).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
