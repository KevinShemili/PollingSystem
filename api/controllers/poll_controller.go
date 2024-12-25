package controllers

import (
	"gin/api/requests"
	"gin/application/usecase/poll/commands/contracts"
	"gin/application/utility"
	"gin/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PollController struct {
	CreatePollCommand  contracts.ICreatePollCommand
	AddVoteCommand     contracts.IAddVoteCommand
	DeletePollCommmand contracts.IDeletePollCommand
	EndPollCommand     contracts.IEndPollCommand
}

func NewPollController(
	CreatePollCommand contracts.ICreatePollCommand,
	AddVoteCommand contracts.IAddVoteCommand,
	DeletePollCommand contracts.IDeletePollCommand,
	EndPollCommand contracts.IEndPollCommand) *PollController {
	return &PollController{
		CreatePollCommand:  CreatePollCommand,
		AddVoteCommand:     AddVoteCommand,
		DeletePollCommmand: DeletePollCommand,
		EndPollCommand:     EndPollCommand}
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

// AddVote handles voting on a specific poll.
//
// @Summary Vote on a poll
// @Description Add a vote to a specific poll category by providing the poll ID in the route and the category ID in the request body. The user must be authenticated.
// @Tags Polls
// @Accept json
// @Produce json
// @Param id path int true "Poll ID"
// @Param request body requests.AddVoteRequest true "Add Vote Request"
// @Success 200 {object} bool "Vote added successfully"
// @Failure 400 {object} utility.ErrorCode "Bad Request - Invalid input"
// @Failure 401 {object} utility.ErrorCode "Unauthorized - Invalid or missing token"
// @Failure 404 {object} utility.ErrorCode "Poll or category not found"
// @Failure 500 {object} utility.ErrorCode "Internal server error"
// @Router /polls/{id}/vote [post]
// @Security BearerAuth
func (uc *PollController) AddVote(c *gin.Context) {

	var request requests.AddVoteRequest

	if err := c.Bind(&request); err != nil {
		c.JSON(utility.BindFailure.StatusCode, utility.BindFailure)
		return
	}

	pollIDString := c.Param("id")
	pollID, errParse := strconv.ParseUint(pollIDString, 10, 32)
	if errParse != nil {
		c.JSON(utility.RouteParameterCast.StatusCode, utility.RouteParameterCast)
		return
	}
	request.PollID = uint(pollID)

	userAny, ok := c.Get("user")
	if !ok {
		c.JSON(utility.Unauthorized.StatusCode, utility.Unauthorized)
	}

	user, ok := userAny.(*entities.User)
	if !ok {
		c.JSON(utility.InternalServerError.StatusCode, utility.InternalServerError)
	}

	result, err := uc.AddVoteCommand.AddVote(&request, user)

	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// DeletePoll handles deleting a specific poll.
//
// @Summary Delete a poll
// @Description Delete a poll by providing the poll ID in the route. The user must be authenticated.
// @Tags Polls
// @Accept json
// @Produce json
// @Param id path int true "Poll ID"
// @Success 200 {object} bool "Poll deleted successfully"
// @Failure 400 {object} utility.ErrorCode "Bad Request - Invalid input"
// @Failure 401 {object} utility.ErrorCode "Unauthorized - Invalid or missing token"
// @Failure 404 {object} utility.ErrorCode "Poll not found"
// @Failure 500 {object} utility.ErrorCode "Internal server error"
// @Router /polls/{id} [delete]
// @Security BearerAuth
func (uc *PollController) DeletePoll(c *gin.Context) {

	pollIDString := c.Param("id")
	pollID, errParse := strconv.ParseUint(pollIDString, 10, 32)
	if errParse != nil {
		c.JSON(utility.RouteParameterCast.StatusCode, utility.RouteParameterCast)
		return
	}

	result, err := uc.DeletePollCommmand.DeletePoll(uint(pollID))

	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// EndPoll handles ending a specific poll.
//
// @Summary End a poll
// @Description End a poll by providing the poll ID in the route. The user must be authenticated.
// @Tags Polls
// @Accept json
// @Produce json
// @Param id path int true "Poll ID"
// @Success 200 {object} bool "Poll ended successfully"
// @Failure 400 {object} utility.ErrorCode "Bad Request - Invalid input"
// @Failure 401 {object} utility.ErrorCode "Unauthorized - Invalid or missing token"
// @Failure 404 {object} utility.ErrorCode "Poll not found"
// @Failure 500 {object} utility.ErrorCode "Internal server error"
// @Router /polls/{id}/end [put]
// @Security BearerAuth
func (uc *PollController) EndPoll(c *gin.Context) {

	pollIDString := c.Param("id")
	pollID, errParse := strconv.ParseUint(pollIDString, 10, 32)
	if errParse != nil {
		c.JSON(utility.RouteParameterCast.StatusCode, utility.RouteParameterCast)
		return
	}

	result, err := uc.EndPollCommand.EndPoll(uint(pollID))

	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, result)
}
