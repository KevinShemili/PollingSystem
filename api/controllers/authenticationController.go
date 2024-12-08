package controllers

import (
	"gin/application/usecase/authentication/commands"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthenticationController struct {
	registerUserCommand commands.IRegisterUserCommand
}

func NewAuthenticationController(registerUserCommand commands.IRegisterUserCommand) *AuthenticationController {
	return &AuthenticationController{registerUserCommand: registerUserCommand}
}

func (uc *AuthenticationController) Register(c *gin.Context) {

	success, err := uc.registerUserCommand.Register()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": success})
}

func (uc *AuthenticationController) Login(c *gin.Context) {

	success, err := uc.registerUserCommand.Register()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": success})
}

func (uc *AuthenticationController) LogOut(c *gin.Context) {

	success, err := uc.registerUserCommand.Register()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": success})
}
