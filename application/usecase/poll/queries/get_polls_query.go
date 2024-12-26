package queries

import (
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

func (r GetPollsQuery) GetPolls(parameters utility.QueryParams) (utility.PaginatedResponse[results.GetPollResult], *utility.ErrorCode) {

	polls, err := r.UnitOfWork.IPollRepository().GetPollsPaginated(parameters)
	if err != nil {
		return utility.PaginatedResponse[results.GetPollResult]{}, utility.InternalServerError.WithDescription(err.Error())
	}

	results := utility.MapPoll(polls)

	return results, nil
}
