package commands

import (
	"gin/api/requests"
	"gin/application/repository/contracts"
	"gin/application/usecase/authentication/results"
	"gin/application/utility"
	"gin/domain/entities"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type LoginCommand struct {
	UnitOfWork contracts.IUnitOfWork
}

func NewLoginCommand(UnitOfWork contracts.IUnitOfWork) *LoginCommand {
	return &LoginCommand{UnitOfWork: UnitOfWork}
}

func (r LoginCommand) Login(request *requests.LoginRequest) (*results.LoginResult, *utility.ErrorCode) {

	uof, err := r.UnitOfWork.Begin()
	if err != nil {
		return nil, utility.InternalServerError.WithDescription(err.Error())
	}
	defer uof.Rollback()

	// check if email exists
	user, err := r.UnitOfWork.Users().GetByEmail(request.Email)

	if err != nil || user == nil {
		return nil, utility.IncorrectEmail
	}

	// verify pass
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(request.Password))
	if err != nil {
		return nil, utility.IncorrectEmail
	}

	// generate token
	signedToken, err := r.generateJWT(user.ID)
	if err != nil {
		return nil, utility.InternalServerError.WithDescription(err.Error())

	}

	// remove old refresh
	oldRefresh, err := r.UnitOfWork.RefreshTokens().GetByUserID(user.ID)
	if err != nil {
		return nil, utility.InternalServerError.WithDescription(err.Error())
	}

	if oldRefresh != nil {
		if err := r.UnitOfWork.RefreshTokens().Delete(oldRefresh.ID); err != nil {
			return nil, utility.InternalServerError.WithDescription(err.Error())
		}
	}

	// get new refresh token & update db
	refreshToken, expiry, err := utility.GenerateRefreshToken()
	if err != nil {
		return nil, utility.InternalServerError.WithDescription(err.Error())
	}

	if err := r.UnitOfWork.RefreshTokens().Create(&entities.RefreshToken{
		Token:     refreshToken,
		Expiry:    expiry,
		JWTToken:  signedToken,
		UserID:    user.ID,
		IsDeleted: false,
	}); err != nil {
		return nil, utility.InternalServerError.WithDescription(err.Error())
	}

	// commit to db
	if err := uof.Commit(); err != nil {
		return nil, utility.InternalServerError.WithDescription(err.Error())
	}

	return &results.LoginResult{
		AuthenticationToken: utility.Encode(signedToken),
		RefreshToken:        utility.Encode(refreshToken),
	}, nil
}

func (r LoginCommand) generateJWT(userID uint) (string, error) {

	jwtExpiryHour, _ := strconv.Atoi(os.Getenv("EXPIRY_JWT"))
	jwtSigningKey := os.Getenv("SECRET_JWT")

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Duration(jwtExpiryHour) * time.Hour).Unix(),
	})

	signedToken, err := jwtToken.SignedString([]byte(jwtSigningKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
