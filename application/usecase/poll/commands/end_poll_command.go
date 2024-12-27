package commands

import (
	"encoding/json"
	repo "gin/application/repository/contracts"
	"gin/application/usecase/poll/commands/contracts"
	"gin/application/usecase/poll/results"
	"gin/application/utility"
	"gin/domain/entities"
	"gin/infrastructure/mail"
	"gin/infrastructure/websocket"
	"log"
)

type EndPollCommand struct {
	UnitOfWork repo.IUnitOfWork
}

func NewEndPollCommand(UnitOfWork repo.IUnitOfWork) contracts.IEndPollCommand {
	return &EndPollCommand{UnitOfWork: UnitOfWork}
}

func (r EndPollCommand) EndPoll(pollID uint, user *entities.User) (bool, *utility.ErrorCode) {

	// begin transaction
	uof, err := r.UnitOfWork.Begin()
	if err != nil {
		return false, utility.InternalServerError.WithDescription(err.Error())
	}
	defer uof.Rollback()

	// check if poll exists
	poll, err := r.UnitOfWork.IPollRepository().GetPollWithVotes(pollID)
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

	// broadcast poll ended
	var broadcastData results.BroadcastExpiry
	broadcastData.BroadcastType = "poll-ended"
	broadcastData.Data.PollID = pollID

	message, _ := json.Marshal(broadcastData)
	websocket.BroadcastMessage(string(message))

	// send email
	go func() {

		results := []map[string]interface{}{}

		for _, category := range poll.Categories {
			results = append(results, map[string]interface{}{
				"CategoryName": category.Name,
				"Votes":        len(category.Votes),
			})
		}

		if err := mail.SendEmail(
			user.Email,
			"Poll Has Ended",
			mail.GetTemplatePath("expired_poll_template.html"),
			map[string]interface{}{
				"PollTitle": poll.Title,
				"Results":   results,
			},
		); err != nil {
			log.Printf("Failed to send email. %v", err)
		}
	}()

	return true, nil
}
