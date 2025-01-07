package contracts

import (
	"gin/application/utility"
	"gin/domain/entities"
)

type IEndPollCommand interface {
	// EndPoll is a method to terminate a poll before automatic expiry
	EndPoll(pollID uint, user *entities.User) (bool, *utility.ErrorCode)
}
