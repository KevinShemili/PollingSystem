package commands

import (
	"gin/api/requests"
	"gin/application/repository/contracts"
	"gin/domain/entities"

	"golang.org/x/crypto/bcrypt"
)

type RegisterCommand struct {
	UserRepository contracts.IUserRepository
}

func NewRegisterCommand(UserRepository contracts.IUserRepository) *RegisterCommand {
	return &RegisterCommand{UserRepository: UserRepository}
}

func (r RegisterCommand) Register(request *requests.RegisterRequest) (bool, error) {

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
