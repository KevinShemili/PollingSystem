package repository

import (
	"errors"
	"gin/application/repository/contracts"

	"gorm.io/gorm"
)

type Repository[T any] struct {
	db *gorm.DB
}

var _ contracts.IRepository[any] = &Repository[any]{}

func NewRepository[T any](db *gorm.DB) *Repository[T] {
	return &Repository[T]{db: db}
}

func (r *Repository[T]) GetAll() ([]T, error) {
	var entities []T
	result := r.db.Where("is_deleted = ?", false).Find(&entities)

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
	result := r.db.Where("id = ? AND is_deleted = ?", id, false).First(&entity)

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

func (r *Repository[T]) Delete(id uint) error {
	var entity T
	result := r.db.First(&entity, id)

	if result.Error != nil {
		return result.Error
	}

	r.db.Model(&entity).Update("is_deleted", true)

	return nil
}
