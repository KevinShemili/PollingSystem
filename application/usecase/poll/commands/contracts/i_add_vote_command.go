package contracts

import (
	"gin/api/requests"
	"gin/application/utility"
	"gin/domain/entities"
)

type IAddVoteCommand interface {
	// AddVote is a method to cast vote on a poll
	AddVote(request *requests.AddVoteRequest, user *entities.User) (bool, *utility.ErrorCode)
}
