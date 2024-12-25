package contracts

import (
	"gin/api/requests"
	"gin/application/usecase/poll/results"
	"gin/application/utility"
)

type IUpdatePollCommand interface {
	UpdatePoll(request *requests.UpdatePollRequest) (*results.UpdatePollResult, *utility.ErrorCode)
}
