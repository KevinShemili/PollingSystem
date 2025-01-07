package contracts

import (
	"gin/api/requests"
	"gin/application/usecase/poll/results"
	"gin/application/utility"
)

type IGetPollsQuery interface {
	// GetPolls returns a list of polls in paginated format
	GetPolls(request *requests.GetPollsRequest) (utility.PaginatedResponse[results.GetPollResult], *utility.ErrorCode)
}
