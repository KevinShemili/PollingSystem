package commands

import (
	"gin/api/requests"
	repo "gin/application/repository/contracts"
	"gin/application/usecase/authentication/commands/contracts"
	"gin/application/utility"

	"github.com/go-playground/validator/v10"
)

type LogOutCommand struct {
	UnitOfWork repo.IUnitOfWork
	Validator  *validator.Validate
}

func NewLogOutCommand(UnitOfWork repo.IUnitOfWork, Validator *validator.Validate) contracts.ILogOutCommand {
	return &LogOutCommand{UnitOfWork: UnitOfWork, Validator: Validator}
}

func (r LogOutCommand) LogOut(request *requests.LogOutRequest) (bool, *utility.ErrorCode) {

	// validate request
	if err := r.Validator.Struct(request); err != nil {
		return false, utility.ValidationError.WithDescription(err.Error())
	}

	// begin transaction
	uof, err := r.UnitOfWork.Begin()
	if err != nil {
		return false, utility.InternalServerError.WithDescription(err.Error())
	}
	defer uof.Rollback()

	// get refresh token
	refreshToken, err := r.UnitOfWork.IRefreshTokenRepository().GetByUserID(uint(request.UserID))
	if err != nil {
		return false, utility.InternalServerError.WithDescription(err.Error())
	}
	if refreshToken == nil {
		return false, utility.Unauthorized
	}

	// delete refresh token
	if err := uof.IRefreshTokenRepository().SoftDelete(refreshToken.ID); err != nil {
		return false, utility.InternalServerError.WithDescription(err.Error())
	}

	// commit to db
	if err := uof.Commit(); err != nil {
		return false, utility.InternalServerError.WithDescription(err.Error())
	}

	return true, nil
}
