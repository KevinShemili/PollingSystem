package commands

import (
	"gin/api/requests"
	"gin/application/repository/contracts"
	"gin/application/usecase/authentication/results"
	"gin/application/utility"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type LoginCommand struct {
	UserRepository contracts.IUserRepository
}

func NewLoginCommand(UserRepository contracts.IUserRepository) *LoginCommand {
	return &LoginCommand{UserRepository: UserRepository}
}

func (r LoginCommand) Login(request *requests.LoginRequest) (*results.LoginResult, *utility.ErrorCode) {

	user, err := r.UserRepository.FindByEmail(request.Email)

	if err != nil || user == nil {
		return nil, utility.IncorrectEmail
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(request.Password))
	if err != nil {
		return nil, utility.IncorrectEmail
	}

	jwtExpiryHour, _ := strconv.Atoi(os.Getenv("EXPIRY_JWT"))
	jwtSigningKey := os.Getenv("SECRET_JWT")

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Duration(jwtExpiryHour) * time.Hour).Unix(),
	})

	signedToken, err := jwtToken.SignedString([]byte(jwtSigningKey))
	if err != nil {
		return nil, utility.InternalServerError.WithDescription(err.Error())
	}

	return &results.LoginResult{
		AuthenticationToken: signedToken,
		RefreshToken:        "",
	}, nil
}
