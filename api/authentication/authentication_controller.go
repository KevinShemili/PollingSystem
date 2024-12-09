package authentication

import (
	"gin/api/requests"
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

// Register handles user registration.
//
// @Summary Register a new user
// @Description This endpoint registers a new user with the provided details.
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body requests.RegisterRequest true "User Registration Request"
// @Success 200 {object} map[string]interface{} "success: true"
// @Failure 400 {object} map[string]interface{} "error: binding error message"
// @Failure 500 {object} map[string]interface{} "error: internal server error"
// @Router /register [post]
// @Security BearerAuth
func (uc *AuthenticationController) Register(c *gin.Context) {

	var request requests.RegisterRequest

	if err := c.Bind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"I AM FAILING BIND": err.Error()})
		return
	}

	success, err := uc.registerUserCommand.Register(&request)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": success})
}

func (uc *AuthenticationController) Login(c *gin.Context) {
}

func (uc *AuthenticationController) LogOut(c *gin.Context) {
}
