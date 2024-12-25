package contracts

import (
	"gin/application/utility"
)

type IDeletePollCommand interface {
	DeletePoll(pollID uint) (bool, *utility.ErrorCode)
}
