package commands

import (
	"gin/api/requests"
	repo "gin/application/repository/contracts"
	"gin/application/usecase/authentication/commands/contracts"
	"gin/application/usecase/authentication/results"
	"gin/application/utility"
	"gin/domain/entities"

	"golang.org/x/crypto/bcrypt"
)

type LoginCommand struct {
	UnitOfWork repo.IUnitOfWork
}

func NewLoginCommand(UnitOfWork repo.IUnitOfWork) contracts.ILoginCommand {
	return &LoginCommand{UnitOfWork: UnitOfWork}
}

func (r LoginCommand) Login(request *requests.LoginRequest) (*results.LoginResult, *utility.ErrorCode) {

	uof, err := r.UnitOfWork.Begin()
	if err != nil {
		return nil, utility.InternalServerError.WithDescription(err.Error())
	}
	defer uof.Rollback()

	// check if email exists
	user, err := r.UnitOfWork.IUserRepository().GetByEmail(request.Email)

	if err != nil || user == nil {
		return nil, utility.IncorrectEmail
	}

	// verify pass
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(request.Password))
	if err != nil {
		return nil, utility.IncorrectEmail
	}

	// generate token
	signedToken, err := utility.GenerateJWTWithID(user.ID)
	if err != nil {
		return nil, utility.InternalServerError.WithDescription(err.Error())

	}

	// remove old refresh
	oldRefresh, err := r.UnitOfWork.IRefreshTokenRepository().GetByUserID(user.ID)
	if err != nil {
		return nil, utility.InternalServerError.WithDescription(err.Error())
	}

	if oldRefresh != nil {
		if err := r.UnitOfWork.IRefreshTokenRepository().SoftDelete(oldRefresh.ID); err != nil {
			return nil, utility.InternalServerError.WithDescription(err.Error())
		}
	}

	// get new refresh token & update db
	refreshToken, expiry, err := utility.GenerateRefreshToken()
	if err != nil {
		return nil, utility.InternalServerError.WithDescription(err.Error())
	}

	if err := r.UnitOfWork.IRefreshTokenRepository().Create(&entities.RefreshToken{
		Token:    refreshToken,
		Expiry:   expiry,
		JWTToken: signedToken,
		UserID:   user.ID,
	}); err != nil {
		return nil, utility.InternalServerError.WithDescription(err.Error())
	}

	// commit to db
	if err := uof.Commit(); err != nil {
		return nil, utility.InternalServerError.WithDescription(err.Error())
	}

	return &results.LoginResult{
		JWTToken:     signedToken,
		RefreshToken: utility.Encode(refreshToken),
	}, nil
}
