package commands

import (
	"encoding/json"
	repo "gin/application/repository/contracts"
	"gin/application/usecase/poll/commands/contracts"
	"gin/application/usecase/poll/results"
	"gin/application/utility"
	"gin/domain/entities"
	"gin/infrastructure/websocket"
)

type DeletePollCommand struct {
	UnitOfWork repo.IUnitOfWork
}

func NewDeletePollCommand(UnitOfWork repo.IUnitOfWork) contracts.IDeletePollCommand {
	return &DeletePollCommand{UnitOfWork: UnitOfWork}
}

func (r DeletePollCommand) DeletePoll(pollID uint, user *entities.User) (bool, *utility.ErrorCode) {

	// begin transaction
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

	// check if user is the creator
	if poll.CreatorID != user.ID {
		return false, utility.NotPollOwner
	}

	// delete poll
	if err := r.UnitOfWork.IPollRepository().SoftDelete(pollID); err != nil {
		return false, utility.InternalServerError.WithDescription(err.Error())
	}

	// commit to db
	if err := uof.Commit(); err != nil {
		return false, utility.InternalServerError.WithDescription(err.Error())
	}

	// broadcast poll deletion
	var broadcastData results.BroadcastDeletion
	broadcastData.BroadcastType = "poll-deleted"
	broadcastData.Data.PollID = pollID

	message, _ := json.Marshal(broadcastData)
	websocket.BroadcastMessage(string(message))

	return true, nil
}
