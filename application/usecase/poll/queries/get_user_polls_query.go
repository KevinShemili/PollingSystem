package queries

import (
	"gin/api/requests"
	repo "gin/application/repository/contracts"
	"gin/application/usecase/poll/queries/contracts"
	"gin/application/usecase/poll/results"
	"gin/application/utility"
)

type GetUserPollsQuery struct {
	UnitOfWork repo.IUnitOfWork
}

func NewGetUserPollsQuery(UnitOfWork repo.IUnitOfWork) contracts.IGetUserPollsQuery {
	return &GetUserPollsQuery{UnitOfWork: UnitOfWork}
}

func (r GetUserPollsQuery) GetPolls(userID uint, request *requests.GetPollsRequest) (utility.PaginatedResponse[results.GetPollResult], *utility.ErrorCode) {

	polls, err := r.UnitOfWork.IPollRepository().GetPollsByUserPaginated(userID, request.QueryParams, request.ShowActiveOnly)
	if err != nil {
		return utility.PaginatedResponse[results.GetPollResult]{}, utility.InternalServerError.WithDescription(err.Error())
	}

	results := utility.MapPoll(polls)

	return results, nil
}
