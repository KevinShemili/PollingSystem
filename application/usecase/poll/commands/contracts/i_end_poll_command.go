package contracts

import (
	"gin/application/utility"
)

type IEndPollCommand interface {
	EndPoll(pollID uint) (bool, *utility.ErrorCode)
}
