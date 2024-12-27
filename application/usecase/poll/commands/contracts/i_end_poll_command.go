package contracts

import (
	"gin/application/utility"
	"gin/domain/entities"
)

type IEndPollCommand interface {
	EndPoll(pollID uint, user *entities.User) (bool, *utility.ErrorCode)
}
