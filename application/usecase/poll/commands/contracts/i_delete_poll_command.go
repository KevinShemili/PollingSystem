package contracts

import (
	"gin/application/utility"
	"gin/domain/entities"
)

type IDeletePollCommand interface {
	DeletePoll(pollID uint, user *entities.User) (bool, *utility.ErrorCode)
}
