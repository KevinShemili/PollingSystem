package contracts

import (
	"gin/application/usecase/poll/results"
	"gin/application/utility"
)

type IGetPollQuery interface {
	// GetPoll returns a poll by its ID
	GetPoll(pollID uint) (*results.GetPollResult, *utility.ErrorCode)
}
