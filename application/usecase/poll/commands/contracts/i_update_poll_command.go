package contracts

import (
	"gin/api/requests"
	"gin/application/utility"
)

type IUpdatePollCommand interface {
	UpdatePoll(userID uint, request *requests.UpdatePollRequest) (bool, *utility.ErrorCode)
}
