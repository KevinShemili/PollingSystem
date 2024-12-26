package contracts

import (
	"gin/application/usecase/poll/results"
	"gin/application/utility"
)

type IGetPollQuery interface {
	GetPoll(pollID uint) (*results.GetPollResult, *utility.ErrorCode)
}
