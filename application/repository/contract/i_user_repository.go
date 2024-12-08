package contract

import "gin/domain/entities"

type IUserRepository interface {
	IRepository[entities.User]

	FindByEmail(email string) (string, error)
}
