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

// Register handles user registration.
//
// @Summary Register a new user
// @Description This endpoint registers a new user with the provided details.
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body requests.RegisterRequest true "User Registration Request"
// @Success 200 {object} map[string]interface{} "success: true"
// @Failure 400 {object} utility.ErrorCode "Binding failure or validation errors"
// @Failure 500 {object} utility.ErrorCode "Internal server error"
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

// Login handles user login.
//
// @Summary Login a user
// @Description Authenticate a user with email and password, returning a JWT token.
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body requests.LoginRequest true "Login Request"
// @Success 200 {object} results.LoginResult "Authentication token and refresh token"
// @Failure 400 {object} utility.ErrorCode "Binding failure"
// @Failure 401 {object} utility.ErrorCode "Invalid credentials"
// @Failure 500 {object} utility.ErrorCode "Internal server error"
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

// Refresh handles token refreshing.
//
// @Summary Refresh user tokens
// @Description Refresh access and refresh tokens using the provided tokens.
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body requests.TokensRequest true "Token Refresh Request"
// @Success 200 {object} results.RefreshResult "New access and refresh tokens"
// @Failure 400 {object} utility.ErrorCode "Binding failure"
// @Failure 401 {object} utility.ErrorCode "Unauthorized"
// @Failure 500 {object} utility.ErrorCode "Internal server error"
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

// LogOut handles user logout.
//
// @Summary Log out a user
// @Description Ends the user session by invalidating the token (requires JWT).
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body requests.LogOutRequest false "LogOut Request (optional)"
// @Success 200 {string} string "Successfully logged out"
// @Failure 401 {object} utility.ErrorCode "Unauthorized - Invalid or missing token"
// @Failure 500 {object} utility.ErrorCode "Internal server error"
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

	c.JSON(http.StatusOK, result)
}
