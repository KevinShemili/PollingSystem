package queries

import (
	repo "gin/application/repository/contracts"
	"gin/application/usecase/poll/queries/contracts"
	"gin/application/usecase/poll/results"
	"gin/application/utility"
)

type GetPollQuery struct {
	UnitOfWork repo.IUnitOfWork
}

func NewGetPollQuery(UnitOfWork repo.IUnitOfWork) contracts.IGetPollQuery {
	return &GetPollQuery{UnitOfWork: UnitOfWork}
}

func (r GetPollQuery) GetPoll(pollID uint) (*results.GetPollResult, *utility.ErrorCode) {

	// get poll
	poll, err := r.UnitOfWork.IPollRepository().GetPollWithVotes(pollID)
	if err != nil {
		return nil, utility.InternalServerError.WithDescription(err.Error())
	}
	if poll == nil {
		return nil, utility.InvalidPollID
	}

	// map result
	result := utility.MapSinglePoll(poll)

	return result, nil
}
