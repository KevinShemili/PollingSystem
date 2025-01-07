package contracts

import (
	"gin/api/requests"
	"gin/application/usecase/poll/results"
	"gin/application/utility"
)

type IGetUserPollsQuery interface {
	// GetPolls returns a list of polls created by a given user, in paginated format 
	GetPolls(userID uint, request *requests.GetPollsRequest) (utility.PaginatedResponse[results.GetPollResult], *utility.ErrorCode)
}
