package contracts

import "gin/domain/entities"

type IUserRepository interface {
	IRepository[entities.User]

	FindByEmail(email string) (*entities.User, error)
}
