package contract

type IRepository[T any] interface {
	Get(id int) (*T, error)
	Create(entity *T, persist bool) error
	Update(entity *T, persist bool) error
	Delete(id int, persist bool) error
	SaveChanges() error
}
