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
	"time"
)

type CreatePollCommand struct {
	UnitOfWork repo.IUnitOfWork
}

func NewCreatePollCommand(UnitOfWork repo.IUnitOfWork) contracts.ICreatePollCommand {
	return &CreatePollCommand{UnitOfWork: UnitOfWork}
}

func (r CreatePollCommand) CreatePoll(request *requests.CreatePollRequest, user *entities.User) (*results.CreatePollResult, *utility.ErrorCode) {

	uof, err := r.UnitOfWork.Begin()
	if err != nil {
		return nil, utility.InternalServerError.WithDescription(err.Error())
	}
	defer uof.Rollback()

	// Convert to time.Time
	expiresAt, err := time.Parse(time.RFC3339, request.ExpiresAt)
	if err != nil {
		return nil, utility.InternalServerError.WithDescription(err.Error())
	}

	var categories []entities.PollCategory

	for _, catName := range request.Categories {
		categories = append(categories, entities.PollCategory{
			Name: catName,
		})
	}

	pollEntity := &entities.Poll{
		Title:      request.Title,
		ExpiresAt:  expiresAt,
		IsEnded:    false,
		CreatorID:  user.ID,
		Categories: categories,
	}

	if err := r.UnitOfWork.IPollRepository().Create(pollEntity); err != nil {
		return nil, utility.InternalServerError.WithDescription(err.Error())
	}

	// commit to db
	if err := uof.Commit(); err != nil {
		return nil, utility.InternalServerError.WithDescription(err.Error())
	}

	var broadcastData results.BroadcastPoll
	broadcastData.BroadcastType = "new-poll"
	broadcastData.Data.PollID = pollEntity.ID
	broadcastData.Data.Title = pollEntity.Title
	broadcastData.Data.ExpiresAt = pollEntity.ExpiresAt
	broadcastData.Data.Ended = pollEntity.IsEnded

	for _, category := range pollEntity.Categories {
		broadcastData.Data.Categories = append(broadcastData.Data.Categories, struct {
			ID    uint   `json:"category_id"`
			Name  string `json:"category_name"`
			Votes int    `json:"category_votes"`
		}{
			ID:    category.ID,
			Name:  category.Name,
			Votes: 0, // just created, 0 votes
		})
	}

	message, _ := json.Marshal(broadcastData)
	websocket.BroadcastMessage(string(message))

	return &results.CreatePollResult{
		Title:      pollEntity.Title,
		ExpiresAt:  pollEntity.ExpiresAt,
		IsEnded:    false,
		Categories: request.Categories,
	}, nil
}
