package contracts

import (
	"gin/api/requests"
	"gin/application/utility"
	"gin/domain/entities"
)

type IAddVoteCommand interface {
	AddVote(request *requests.AddVoteRequest, user *entities.User) (bool, *utility.ErrorCode)
}
