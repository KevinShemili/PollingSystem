package contracts

import (
	"gin/api/requests"
	"gin/application/usecase/poll/results"
	"gin/application/utility"
)

type IGetPollsQuery interface {
	GetPolls(request *requests.GetPollsRequest) (utility.PaginatedResponse[results.GetPollResult], *utility.ErrorCode)
}
