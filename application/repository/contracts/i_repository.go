package contracts

type IRepository[T any] interface {
	GetByID(id uint) (*T, error)
	GetAll() ([]T, error)
	Create(entity *T) error
	Update(entity *T) error
	Delete(id uint) error
}
