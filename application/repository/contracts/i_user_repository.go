package contracts

import "gin/domain/entities"

type IUserRepository interface {
	IRepository[entities.User]

	GetByEmail(email string) (*entities.User, error)
}
