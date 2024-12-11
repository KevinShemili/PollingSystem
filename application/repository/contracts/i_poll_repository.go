package contracts

import "gin/domain/entities"

type IPollRepository interface {
	IRepository[entities.Poll]
}
