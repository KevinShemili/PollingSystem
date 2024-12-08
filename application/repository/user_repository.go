package repository

import (
	"gin/domain/entities"

	"gorm.io/gorm"
)

type UserRepository struct {
	*Repository[entities.User]
}

// NewUserRepository creates a new instance of UserRepository.
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		Repository: NewRepository[entities.User](db),
	}
}

// Implement additional methods for IUserRepository
func (r *UserRepository) FindByEmail(email string) (string, error) {
	return "I HAVE REACHED REPO", nil

	/*var user entities.User
	result := r.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil*/
}
