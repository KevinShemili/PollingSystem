package commands

import (
	"gin/api/requests"
	repo "gin/application/repository/contracts"
	"gin/application/usecase/authentication/commands/contracts"
	"gin/application/utility"
	"gin/domain/entities"

	"golang.org/x/crypto/bcrypt"
)

type RegisterCommand struct {
	UnitOfWork repo.IUnitOfWork
}

func NewRegisterCommand(UnitOfWork repo.IUnitOfWork) contracts.IRegisterCommand {
	return &RegisterCommand{UnitOfWork: UnitOfWork}
}

func (r RegisterCommand) Register(request *requests.RegisterRequest) (bool, *utility.ErrorCode) {

	duplicate, err := r.UnitOfWork.IUserRepository().GetByEmail(request.Email)
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

	if err := r.UnitOfWork.IUserRepository().Create(&entities.User{
		Email:        request.Email,
		PasswordHash: string(hash),
	}); err != nil {
		return false, utility.InternalServerError.WithDescription(err.Error())
	}

	return true, nil
}
