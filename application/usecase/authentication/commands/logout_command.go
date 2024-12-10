package commands

import (
	"gin/api/requests"
	"gin/application/repository/contracts"
	"gin/application/utility"
)

type LogOutCommand struct {
	UnitOfWork contracts.IUnitOfWork
}

func NewLogOutCommand(UnitOfWork contracts.IUnitOfWork) *LogOutCommand {
	return &LogOutCommand{UnitOfWork: UnitOfWork}
}

func (r LogOutCommand) LogOut(request *requests.LogOutRequest) (bool, *utility.ErrorCode) {

	uof, err := r.UnitOfWork.Begin()
	if err != nil {
		return false, utility.InternalServerError.WithDescription(err.Error())
	}
	defer uof.Rollback()

	refreshToken, err := r.UnitOfWork.RefreshTokens().GetByUserID(uint(request.UserID))
	if err != nil {
		return false, utility.InternalServerError.WithDescription(err.Error())
	}

	if err := uof.RefreshTokens().Delete(refreshToken.ID); err != nil {
		return false, utility.InternalServerError.WithDescription(err.Error())
	}

	// commit to db
	if err := uof.Commit(); err != nil {
		return false, utility.InternalServerError.WithDescription(err.Error())
	}

	return true, nil
}
