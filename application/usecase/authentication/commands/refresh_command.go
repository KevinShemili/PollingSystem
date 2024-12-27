package commands

import (
	"errors"
	"gin/api/requests"
	repo "gin/application/repository/contracts"
	"gin/application/usecase/authentication/commands/contracts"
	"gin/application/usecase/authentication/results"
	"gin/application/utility"
	"gin/domain/entities"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type RefreshCommand struct {
	UnitOfWork repo.IUnitOfWork
	Validator  *validator.Validate
}

func NewRefreshCommand(UnitOfWork repo.IUnitOfWork, Validator *validator.Validate) contracts.IRefreshCommand {
	return &RefreshCommand{UnitOfWork: UnitOfWork, Validator: Validator}
}

func (r RefreshCommand) Refresh(request *requests.TokensRequest) (*results.RefreshResult, *utility.ErrorCode) {

	if err := r.Validator.Struct(request); err != nil {
		return nil, utility.ValidationError.WithDescription(err.Error())
	}

	uof, err := r.UnitOfWork.Begin()
	if err != nil {
		return nil, utility.InternalServerError.WithDescription(err.Error())
	}
	defer uof.Rollback()

	// decode refresh
	decodedRefresh, err := utility.Decode(request.RefreshToken)
	if err != nil {
		return nil, utility.InvalidToken.WithDescription(err.Error())
	}

	// get claims from JWT
	claims, err := r.parseClaims(request.JWTToken)
	if err != nil {
		return nil, utility.InvalidToken.WithDescription(err.Error())
	}

	// check if user exists
	userID := int(claims["sub"].(float64))
	user, err := r.UnitOfWork.IUserRepository().GetByID(uint(userID))
	if err != nil {
		return nil, utility.InternalServerError.WithDescription(err.Error())
	}
	if user == nil {
		return nil, utility.Unauthorized
	}

	// check if refresh is correct, JWT is correct & refresh is NOT expired
	currentRefresh, err := r.UnitOfWork.IRefreshTokenRepository().GetByUserID(user.ID)
	if err != nil {
		return nil, utility.InternalServerError.WithDescription(err.Error())
	}
	if currentRefresh.Token != decodedRefresh || currentRefresh.JWTToken != request.JWTToken ||
		currentRefresh.Expiry.Before(time.Now()) {
		return nil, utility.Unauthorized
	}

	// Generate a new jwt & refresh
	newJWT, err := utility.GenerateJWTWithClaims(claims)
	if err != nil {
		return nil, utility.InternalServerError.WithDescription(err.Error())
	}

	newRefresh, expiry, err := utility.GenerateRefreshToken()
	if err != nil {
		return nil, utility.InternalServerError.WithDescription(err.Error())
	}

	// Delete old refresh
	if err := uof.IRefreshTokenRepository().SoftDelete(currentRefresh.ID); err != nil {
		return nil, utility.InternalServerError.WithDescription(err.Error())
	}

	// Save new refresh
	if err := uof.IRefreshTokenRepository().Create(&entities.RefreshToken{
		Token:    newRefresh,
		Expiry:   expiry,
		JWTToken: newJWT,
		UserID:   user.ID,
	}); err != nil {
		return nil, utility.InternalServerError.WithDescription(err.Error())
	}

	// commit
	if err := uof.Commit(); err != nil {
		return nil, utility.InternalServerError.WithDescription(err.Error())
	}

	return &results.RefreshResult{
		JWTToken:     newJWT,
		RefreshToken: utility.Encode(newRefresh),
	}, nil
}

func (r RefreshCommand) parseClaims(jwtString string) (jwt.MapClaims, error) {
	jwtSigningKey := os.Getenv("SECRET_JWT")

	jwtToken, err := jwt.Parse(jwtString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(jwtSigningKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok || !jwtToken.Valid {
		return nil, errors.New("JWT is not valid")
	}

	return claims, nil
}
