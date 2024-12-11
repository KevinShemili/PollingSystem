package controllers

import (
	"gin/api/requests"
	"gin/application/usecase/poll/commands/contracts"
	"gin/application/utility"
	"gin/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PollController struct {
	CreatePollCommand contracts.ICreatePollCommand
}

func NewPollController(CreatePollCommand contracts.ICreatePollCommand) *PollController {
	return &PollController{
		CreatePollCommand: CreatePollCommand}
}

// CreatePoll handles poll creation.
//
// @Summary Create a new poll
// @Description Create a new poll with a title, expiration time, and categories. The user must be authenticated.
// @Tags Polls
// @Accept json
// @Produce json
// @Param request body requests.CreatePollRequest true "Create Poll Request"
// @Success 200 {object} results.CreatePollResult "Poll created successfully"
// @Failure 400 {object} utility.ErrorCode "Bad Request - Invalid input"
// @Failure 401 {object} utility.ErrorCode "Unauthorized - Invalid or missing token"
// @Failure 500 {object} utility.ErrorCode "Internal server error"
// @Router /polls [post]
// @Security BearerAuth
func (uc *PollController) CreatePoll(c *gin.Context) {

	var request requests.CreatePollRequest

	if err := c.Bind(&request); err != nil {
		c.JSON(utility.BindFailure.StatusCode, utility.BindFailure)
		return
	}

	userAny, ok := c.Get("user")
	if !ok {
		c.JSON(utility.Unauthorized.StatusCode, utility.Unauthorized)
	}

	user, ok := userAny.(*entities.User)
	if !ok {
		c.JSON(utility.InternalServerError.StatusCode, utility.InternalServerError)
	}

	result, err := uc.CreatePollCommand.CreatePoll(&request, user)

	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, result)
}
