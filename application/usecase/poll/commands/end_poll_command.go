package commands

import (
	repo "gin/application/repository/contracts"
	"gin/application/usecase/poll/commands/contracts"
	"gin/application/utility"
)

type EndPollCommand struct {
	UnitOfWork repo.IUnitOfWork
}

func NewEndPollCommand(UnitOfWork repo.IUnitOfWork) contracts.IEndPollCommand {
	return &EndPollCommand{UnitOfWork: UnitOfWork}
}

func (r EndPollCommand) EndPoll(pollID uint) (bool, *utility.ErrorCode) {

	uof, err := r.UnitOfWork.Begin()
	if err != nil {
		return false, utility.InternalServerError.WithDescription(err.Error())
	}
	defer uof.Rollback()

	// check if poll exists
	poll, err := r.UnitOfWork.IPollRepository().GetByID(pollID)
	if err != nil {
		return false, utility.InternalServerError.WithDescription(err.Error())
	}
	if poll == nil {
		return false, utility.InvalidPollID
	}

	// check if already ended
	if poll.IsEnded {
		return false, utility.AlreadyEnded
	}

	// end poll
	poll.IsEnded = true
	if err := r.UnitOfWork.IPollRepository().Update(poll); err != nil {
		return false, utility.InternalServerError.WithDescription(err.Error())
	}

	// commit to db
	if err := uof.Commit(); err != nil {
		return false, utility.InternalServerError.WithDescription(err.Error())
	}

	return true, nil
}
