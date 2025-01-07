package contracts

import "gin/domain/entities"

type IPollCategoryRepository interface {
	IRepository[entities.PollCategory]

	// HasCategory checks if a poll has a given category ID
	HasCategory(PollID uint, PollCategoryID uint) (bool, error)
}
