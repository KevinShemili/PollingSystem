package contracts

import (
	"gin/api/requests"
	"gin/application/usecase/poll/results"
	"gin/application/utility"
	"gin/domain/entities"
)

type ICreatePollCommand interface {
	CreatePoll(request *requests.CreatePollRequest, user *entities.User) (*results.CreatePollResult, *utility.ErrorCode)
}
