package repository

import (
	"errors"

	"gorm.io/gorm"
)

type Repository[T any] struct {
	db *gorm.DB
}

func NewRepository[T any](db *gorm.DB) *Repository[T] {
	return &Repository[T]{db: db}
}

func (r *Repository[T]) GetAll() ([]T, error) {

	var entities []T

	result := r.db.Find(&entities)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, result.Error
	}

	return entities, nil
}

func (r *Repository[T]) GetByID(id uint) (*T, error) {
	var entity T
	result := r.db.First(&entity, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, result.Error
	}

	return &entity, nil
}

func (r *Repository[T]) Create(entity *T) error {
	return r.db.Create(entity).Error
}

func (r *Repository[T]) Update(entity *T) error {
	return r.db.Save(entity).Error
}

func (r *Repository[T]) SoftDelete(id uint) error {

	// take it for granted there is no error
	// logic checks are done in upper layers
	entity, _ := r.GetByID(id)

	result := r.db.Delete(entity)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *Repository[T]) HardDelete(id uint) error {

	// take it for granted there is no error
	// logic checks are done in upper layers
	entity, _ := r.GetByID(id)

	result := r.db.
		Unscoped().
		Delete(entity)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
