package contracts

import (
	"gin/application/usecase/poll/results"
	"gin/application/utility"
)

type IGetUserPollsQuery interface {
	GetPolls(userID uint, parameters utility.QueryParams) (utility.PaginatedResponse[results.GetPollResult], *utility.ErrorCode)
}
