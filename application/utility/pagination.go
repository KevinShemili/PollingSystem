package utility

import (
	"fmt"

	"gorm.io/gorm"
)

type QueryParams struct {
	Page     int
	PageSize int
	Filter   string
}

type PaginatedResponse[T any] struct {
	Data       []T   `json:"data"`
	Page       int   `json:"page"`
	PageSize   int   `json:"page_size"`
	TotalCount int64 `json:"total_count"`
	TotalPages int   `json:"total_pages"`
}

// apply filter on title & description
func ApplyFilter(db *gorm.DB, params QueryParams) *gorm.DB {
	if params.Filter != "" {
		filter := fmt.Sprintf("%%%s%%", params.Filter)
		db = db.Where("(title ILIKE ? OR description ILIKE ?)", filter, filter)
	}
	return db
}

func ApplyPagination(db *gorm.DB, params *QueryParams) *gorm.DB {

	if params.Page < 1 {
		params.Page = 1
	}

	if params.PageSize < 1 {
		params.PageSize = 10
	}

	offset := (params.Page - 1) * params.PageSize
	return db.Offset(offset).Limit(params.PageSize)
}

// apply both paginaton & filter
func PaginateAndFilter[T any](db *gorm.DB, params QueryParams) (PaginatedResponse[T], error) {
	var result PaginatedResponse[T]
	var data []T

	query := ApplyFilter(db, params)

	var total int64
	if err := query.Model(new(T)).Count(&total).Error; err != nil {
		return result, err
	}

	query = ApplyPagination(query, &params)

	if err := query.Find(&data).Error; err != nil {
		return result, err
	}

	result.Data = data
	result.Page = params.Page
	result.PageSize = params.PageSize
	result.TotalCount = total
	totalPages := int((total + int64(params.PageSize) - 1) / int64(params.PageSize))
	result.TotalPages = totalPages

	return result, nil
}
