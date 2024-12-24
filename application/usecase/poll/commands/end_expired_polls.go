package commands

import (
	"encoding/json"
	"fmt"
	"gin/application/repository/contracts"
	"gin/application/usecase/poll/results"
	"gin/infrastructure/websocket"
	"time"
)

func EndExpiredPolls(UnitOfWork contracts.IUnitOfWork) error {
	polls, err := UnitOfWork.IPollRepository().GetExpiredPolls(time.Now())
	if err != nil {
		return err
	}

	uof, err := UnitOfWork.Begin()
	if err != nil {
		return err
	}
	defer uof.Rollback()

	for _, poll := range polls {
		poll.IsEnded = true
		if err := uof.IPollRepository().Update(poll); err != nil {
			return err
		}
		fmt.Printf("Ended Poll: %d", poll.ID)

	}

	if err := uof.Commit(); err != nil {
		return err
	}

	for _, endedPoll := range polls {
		var broadcastData results.BroadcastVote
		broadcastData.BroadcastType = "poll-ended"
		broadcastData.Data.PollID = endedPoll.ID

		message, _ := json.Marshal(broadcastData)
		websocket.BroadcastMessage(string(message))
	}
	return nil
}
