package contracts

import "gin/domain/entities"

type IPollCategoryRepository interface {
	IRepository[entities.PollCategory]

	HasCategory(PollID uint, PollCategoryID uint) (bool, error)
}
