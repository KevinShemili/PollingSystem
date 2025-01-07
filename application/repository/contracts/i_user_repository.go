package contracts

import "gin/domain/entities"

type IUserRepository interface {
	IRepository[entities.User]

	// GetByEmail returns the user with the given email, or nil if not found
	GetByEmail(email string) (*entities.User, error)
}
