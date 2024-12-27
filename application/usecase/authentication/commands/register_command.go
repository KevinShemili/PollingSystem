package commands

import (
	"gin/api/requests"
	repo "gin/application/repository/contracts"
	"gin/application/usecase/authentication/commands/contracts"
	"gin/application/utility"
	"gin/domain/entities"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type RegisterCommand struct {
	UnitOfWork repo.IUnitOfWork
	Validator  *validator.Validate
}

func NewRegisterCommand(UnitOfWork repo.IUnitOfWork, Validator *validator.Validate) contracts.IRegisterCommand {
	return &RegisterCommand{UnitOfWork: UnitOfWork, Validator: Validator}
}

func (r RegisterCommand) Register(request *requests.RegisterRequest) (bool, *utility.ErrorCode) {

	if err := r.Validator.Struct(request); err != nil {
		return false, utility.ValidationError.WithDescription(err.Error())
	}

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
