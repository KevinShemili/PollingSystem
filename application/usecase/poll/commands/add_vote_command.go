package commands

import (
	"encoding/json"
	"gin/api/requests"
	repo "gin/application/repository/contracts"
	"gin/application/usecase/poll/commands/contracts"
	"gin/application/usecase/poll/results"
	"gin/application/utility"
	"gin/domain/entities"
	"gin/infrastructure/websocket"
)

type AddVoteCommand struct {
	UnitOfWork repo.IUnitOfWork
}

func NewAddVoteCommand(UnitOfWork repo.IUnitOfWork) contracts.IAddVoteCommand {
	return &AddVoteCommand{UnitOfWork: UnitOfWork}
}

func (r AddVoteCommand) AddVote(request *requests.AddVoteRequest, user *entities.User) (bool, *utility.ErrorCode) {

	uof, err := r.UnitOfWork.Begin()
	if err != nil {
		return false, utility.InternalServerError.WithDescription(err.Error())
	}
	defer uof.Rollback()

	// check if poll exists
	poll, err := r.UnitOfWork.IPollRepository().GetByID(request.PollID)
	if err != nil {
		return false, utility.InternalServerError.WithDescription(err.Error())
	}
	if poll == nil {
		return false, utility.InvalidPollID
	}

	// check if poll has ended
	if poll.IsEnded {
		return false, utility.PollExpired
	}

	// check if poll has category
	hasCategory, err := r.UnitOfWork.IPollCategoryRepository().HasCategory(poll.ID, request.PollCategoryID)
	if err != nil {
		return false, utility.InternalServerError.WithDescription(err.Error())
	}
	if !hasCategory {
		return false, utility.InvalidCategoryID
	}

	// check if already voted
	hasVoted, err := r.UnitOfWork.IVoteRepository().HasAlreadyVoted(request.PollID, user.ID)
	if err != nil {
		return false, utility.InternalServerError.WithDescription(err.Error())
	}
	if hasVoted {
		return false, utility.AlreadyVoted
	}

	vote := &entities.Vote{
		UserID:         user.ID,
		PollCategoryID: request.PollCategoryID,
	}

	// create vote
	if err := r.UnitOfWork.IVoteRepository().Create(vote); err != nil {
		return false, utility.InternalServerError.WithDescription(err.Error())
	}

	// commit to db
	if err := uof.Commit(); err != nil {
		return false, utility.InternalServerError.WithDescription(err.Error())
	}

	// get poll with votes & broadcast
	updatedPoll, err := r.UnitOfWork.IPollRepository().GetPollWithVotes(poll.ID)
	if err != nil {
		return false, utility.InternalServerError.WithDescription(err.Error())
	}

	var broadcastData results.BroadcastVote
	broadcastData.BroadcastType = "update-poll"
	broadcastData.Data.PollID = updatedPoll.ID
	for _, category := range updatedPoll.Categories {
		broadcastData.Data.Categories = append(broadcastData.Data.Categories, struct {
			CategoryID uint `json:"category_id"`
			Votes      int  `json:"category_votes"`
		}{
			CategoryID: category.ID,
			Votes:      len(category.Votes),
		})
	}

	message, _ := json.Marshal(broadcastData)
	websocket.BroadcastMessage(string(message))

	return true, nil
}
