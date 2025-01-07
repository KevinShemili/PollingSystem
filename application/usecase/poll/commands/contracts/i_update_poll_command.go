package contracts

import (
	"gin/api/requests"
	"gin/application/utility"
)

type IUpdatePollCommand interface {
	// UpdatePoll is a method to update a poll
	UpdatePoll(userID uint, request *requests.UpdatePollRequest) (bool, *utility.ErrorCode)
}
