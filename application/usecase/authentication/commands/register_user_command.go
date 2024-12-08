package commands

import (
	"gin/api/requests"
	"gin/application/repository/contract"
	"gin/domain/entities"

	"golang.org/x/crypto/bcrypt"
)

type RegisterUserCommand struct {
	UserRepository contract.IUserRepository
}

func NewRegisterUserCommand(UserRepository contract.IUserRepository) *RegisterUserCommand {
	return &RegisterUserCommand{UserRepository: UserRepository}
}

func (r RegisterUserCommand) Register(request *requests.RegisterRequest) (bool, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	if err != nil {
		print("01")
	}

	err = r.UserRepository.Create(&entities.User{
		Email:        request.Email,
		PasswordHash: string(hash),
	})

	if err != nil {
		return false, err
	}

	return true, nil
}
