package commands

import (
	repo "gin/application/repository/contracts"
	"gin/application/usecase/poll/commands/contracts"
	"gin/application/utility"
)

type DeletePollCommand struct {
	UnitOfWork repo.IUnitOfWork
}

func NewDeletePollCommand(UnitOfWork repo.IUnitOfWork) contracts.IDeletePollCommand {
	return &DeletePollCommand{UnitOfWork: UnitOfWork}
}

func (r DeletePollCommand) DeletePoll(pollID uint) (bool, *utility.ErrorCode) {

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

	// delete poll
	if err := r.UnitOfWork.IPollRepository().Delete(pollID); err != nil {
		return false, utility.InternalServerError.WithDescription(err.Error())
	}

	// commit to db
	if err := uof.Commit(); err != nil {
		return false, utility.InternalServerError.WithDescription(err.Error())
	}

	return true, nil
}
