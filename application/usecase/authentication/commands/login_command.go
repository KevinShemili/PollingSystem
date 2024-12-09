package commands

import (
	"gin/api/requests"
	"gin/application/repository/contracts"
	"gin/application/usecase/authentication/results"
	"gin/application/utility"
)

type LoginCommand struct {
	UserRepository contracts.IUserRepository
}

func NewLoginCommand(UserRepository contracts.IUserRepository) *LoginCommand {
	return &LoginCommand{UserRepository: UserRepository}
}

func (r LoginCommand) Login(request *requests.LoginRequest) (*results.LoginResult, *utility.ErrorCode) {

	_, err := r.UserRepository.FindByEmail(request.Email)

	if err != nil {
		return nil, utility.IncorrectEmail
	}

	return nil, nil
}
