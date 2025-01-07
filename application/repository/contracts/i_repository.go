package contracts

type IRepository[T any] interface {
	// GetByID returns an entity by its ID
	GetByID(id uint) (*T, error)

	// GetAll returns all entities
	GetAll() ([]T, error)

	// Create inserts a new entity
	Create(entity *T) error

	// Update updates an entity
	Update(entity *T) error

	// SoftDelete marks an entity as deleted
	SoftDelete(id uint) error

	// HardDelete deletes an entity from the DB
	HardDelete(id uint) error
}
