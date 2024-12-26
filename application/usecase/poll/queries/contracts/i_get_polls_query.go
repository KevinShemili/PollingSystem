package contracts

import (
	"gin/application/usecase/poll/results"
	"gin/application/utility"
)

type IGetPollsQuery interface {
	GetPolls(parameters utility.QueryParams) (utility.PaginatedResponse[results.GetPollResult], *utility.ErrorCode)
}
