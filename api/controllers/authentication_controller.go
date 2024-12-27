package controllers

import (
	"gin/api/requests"
	"gin/application/usecase/authentication/commands/contracts"
	"gin/application/utility"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthenticationController struct {
	RegisterCommand contracts.IRegisterCommand
	LoginCommand    contracts.ILoginCommand
	RefreshCommand  contracts.IRefreshCommand
	LogOutCommand   contracts.ILogOutCommand
}

func NewAuthenticationController(RegisterCommand contracts.IRegisterCommand,
	LoginCommand contracts.ILoginCommand,
	RefreshCommand contracts.IRefreshCommand,
	LogOutCommand contracts.ILogOutCommand) *AuthenticationController {
	return &AuthenticationController{
		RegisterCommand: RegisterCommand,
		LoginCommand:    LoginCommand,
		RefreshCommand:  RefreshCommand,
		LogOutCommand:   LogOutCommand}
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body requests.RegisterRequest true "User Registration Request"
// @Success 200 {object} bool "success: true"
// @Failure 400 {object} utility.ErrorCode "4xx Errors"
// @Failure 500 {object} utility.ErrorCode "5xx Errors"
// @Router /auth/register [post]
func (uc *AuthenticationController) Register(c *gin.Context) {

	var request requests.RegisterRequest

	if err := c.Bind(&request); err != nil {
		c.JSON(utility.BindFailure.StatusCode, utility.BindFailure)
		return
	}

	result, err := uc.RegisterCommand.Register(&request)

	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// Login godoc
// @Summary Login a user
// @Description Login a user
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body requests.LoginRequest true "Login Request"
// @Success 200 {object} results.LoginResult "JWT Token & Refresh Token"
// @Failure 400 {object} utility.ErrorCode "4xx Errors"
// @Failure 500 {object} utility.ErrorCode "5xx Errors"
// @Router /auth/login [post]
func (uc *AuthenticationController) Login(c *gin.Context) {

	var request requests.LoginRequest

	if err := c.Bind(&request); err != nil {
		c.JSON(utility.BindFailure.StatusCode, utility.BindFailure)
		return
	}

	result, err := uc.LoginCommand.Login(&request)

	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// Refresh godoc
// @Summary Refresh user tokens
// @Description Generates a new JWT token and refresh token using the refresh token provided
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body requests.TokensRequest true "Token Refresh Request"
// @Success 200 {object} results.RefreshResult "JWT & Refresh Token"
// @Failure 400 {object} utility.ErrorCode "4xx Errors"
// @Failure 500 {object} utility.ErrorCode "5xx Errors"
// @Router /auth/refresh [post]
func (uc *AuthenticationController) Refresh(c *gin.Context) {

	var request requests.TokensRequest

	if err := c.Bind(&request); err != nil {
		c.JSON(utility.BindFailure.StatusCode, utility.BindFailure)
		return
	}

	result, err := uc.RefreshCommand.Refresh(&request)

	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// LogOut godoc
// @Summary Log out a user
// @Description Ends the user session by invalidating the token (requires JWT).
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body requests.LogOutRequest true "LogOut Request"
// @Success 200 {object} bool "Successfully logged out"
// @Failure 400 {object} utility.ErrorCode "4xx Errors"
// @Failure 500 {object} utility.ErrorCode "5xx Errors"
// @Router /auth/logout [post]
// @Security BearerAuth
func (uc *AuthenticationController) LogOut(c *gin.Context) {

	var request requests.LogOutRequest

	if err := c.Bind(&request); err != nil {
		c.JSON(utility.BindFailure.StatusCode, utility.BindFailure)
		return
	}

	result, err := uc.LogOutCommand.LogOut(&request)

	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.Set("user", nil)

	c.JSON(http.StatusOK, result)
}
