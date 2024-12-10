package commands

import (
	"gin/api/requests"
	"gin/application/repository/contracts"
	"gin/application/utility"
	"gin/domain/entities"

	"golang.org/x/crypto/bcrypt"
)

type RegisterCommand struct {
	UnitOfWork contracts.IUnitOfWork
}

func NewRegisterCommand(UnitOfWork contracts.IUnitOfWork) *RegisterCommand {
	return &RegisterCommand{UnitOfWork: UnitOfWork}
}

func (r RegisterCommand) Register(request *requests.RegisterRequest) (bool, *utility.ErrorCode) {

	duplicate, err := r.UnitOfWork.Users().GetByEmail(request.Email)
	if err != nil {
		return false, utility.InternalServerError.WithDescription(err.Error())
	}
	if duplicate != nil {
		return false, utility.DuplicateEmail
	}

	if flag := utility.ValidateEmail(request.Email); !flag {
		return false, utility.EmailFormat
	}

	if flag := utility.ValidatePassword(request.Password); !flag {
		return false, utility.PasswordFormat
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	if err != nil {
		return false, utility.InternalServerError.WithDescription(err.Error())
	}

	if err := r.UnitOfWork.Users().Create(&entities.User{
		Email:        request.Email,
		PasswordHash: string(hash),
	}); err != nil {
		return false, utility.InternalServerError.WithDescription(err.Error())
	}

	return true, nil
}
