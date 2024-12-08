package repository

import "gorm.io/gorm"

type Repository[T any] struct {
	db       *gorm.DB
	tx       *gorm.DB // Holds a transaction instance when using Unit of Work
	inTxMode bool     // Tracks if we're in transaction mode for deferred persistence
}

func NewRepository[T any](db *gorm.DB) *Repository[T] {
	return &Repository[T]{db: db, inTxMode: false}
}

func (r *Repository[T]) startTransaction() {
	if !r.inTxMode {
		r.tx = r.db.Begin()
		r.inTxMode = true
	}
}

func (r *Repository[T]) SaveChanges() error {
	if r.inTxMode && r.tx != nil {
		err := r.tx.Commit().Error
		r.inTxMode = false
		r.tx = nil
		return err
	}
	return nil
}

func (r *Repository[T]) Get(id int) (*T, error) {
	var entity T
	err := r.db.First(&entity, id).Error
	return &entity, err
}

func (r *Repository[T]) Create(entity *T, persist bool) error {
	db := r.db
	if !persist {
		r.startTransaction()
		db = r.tx
	}
	return db.Create(entity).Error
}

// Update modifies an existing entity in the database.
func (r *Repository[T]) Update(entity *T, persist bool) error {
	db := r.db
	if !persist {
		r.
		()
		db = r.tx
	}
	return db.Save(entity).Error
}

// Delete removes an entity by ID.
func (r *Repository[T]) Delete(id uint, persist bool) error {
	db := r.db
	if !persist {
		r.startTransaction()
		db = r.tx
	}
	return db.Delete(new(T), id).Error
}
