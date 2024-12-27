package queries

import (
	"gin/api/requests"
	repo "gin/application/repository/contracts"
	"gin/application/usecase/poll/queries/contracts"
	"gin/application/usecase/poll/results"
	"gin/application/utility"
)

type GetPollsQuery struct {
	UnitOfWork repo.IUnitOfWork
}

func NewGetPollsQuery(UnitOfWork repo.IUnitOfWork) contracts.IGetPollsQuery {
	return &GetPollsQuery{UnitOfWork: UnitOfWork}
}

func (r GetPollsQuery) GetPolls(request *requests.GetPollsRequest) (utility.PaginatedResponse[results.GetPollResult], *utility.ErrorCode) {

	polls, err := r.UnitOfWork.IPollRepository().GetPollsPaginated(request.QueryParams, request.ShowActiveOnly)
	if err != nil {
		return utility.PaginatedResponse[results.GetPollResult]{}, utility.InternalServerError.WithDescription(err.Error())
	}

	results := utility.MapPoll(polls)

	return results, nil
}
