package commands

import (
	"gin/application/repository/contract"
)

type RegisterUserCommand struct {
	userRepository contract.IUserRepository
}

func NewRegisterUserCommand(userRepository contract.IUserRepository) *RegisterUserCommand {
	return &RegisterUserCommand{userRepository: userRepository}
}

func (r RegisterUserCommand) Register() (string, error) {

	return r.userRepository.FindByEmail("test")
}
