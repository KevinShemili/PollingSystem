package contracts

import (
	"gin/application/utility"
	"gin/domain/entities"
)

type IDeletePollCommand interface {
	// DeletePoll is a method to soft delete a poll
	DeletePoll(pollID uint, user *entities.User) (bool, *utility.ErrorCode)
}
