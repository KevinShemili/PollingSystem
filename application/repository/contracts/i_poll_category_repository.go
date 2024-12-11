package contracts

import "gin/domain/entities"

type IPollCategoryRepository interface {
	IRepository[entities.PollCategory]
}
