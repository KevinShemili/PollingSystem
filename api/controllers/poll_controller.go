package controllers

import (
	"gin/api/requests"
	commands "gin/application/usecase/poll/commands/contracts"
	queries "gin/application/usecase/poll/queries/contracts"
	"gin/application/utility"
	"gin/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PollController struct {
	CreatePollCommand  commands.ICreatePollCommand
	AddVoteCommand     commands.IAddVoteCommand
	DeletePollCommmand commands.IDeletePollCommand
	EndPollCommand     commands.IEndPollCommand
	UpdatePollCommand  commands.IUpdatePollCommand
	GetPollQuery       queries.IGetPollQuery
	GetPollsQuery      queries.IGetPollsQuery
	GetUserPollsQuery  queries.IGetUserPollsQuery
}

func NewPollController(
	CreatePollCommand commands.ICreatePollCommand,
	AddVoteCommand commands.IAddVoteCommand,
	DeletePollCommand commands.IDeletePollCommand,
	EndPollCommand commands.IEndPollCommand,
	GetPollQuery queries.IGetPollQuery,
	GetPollsQuery queries.IGetPollsQuery,
	GetUserPollsQuery queries.IGetUserPollsQuery,
	UpdatePollCommand commands.IUpdatePollCommand) *PollController {
	return &PollController{
		CreatePollCommand:  CreatePollCommand,
		AddVoteCommand:     AddVoteCommand,
		DeletePollCommmand: DeletePollCommand,
		EndPollCommand:     EndPollCommand,
		GetPollQuery:       GetPollQuery,
		GetPollsQuery:      GetPollsQuery,
		GetUserPollsQuery:  GetUserPollsQuery,
		UpdatePollCommand:  UpdatePollCommand}
}

// CreatePoll godoc
// @Summary Create a new poll
// @Description Create a new poll in the system.
// @Tags Polls
// @Accept json
// @Produce json
// @Param request body requests.CreatePollRequest true "Create Poll Request"
// @Success 200 {object} results.CreatePollResult "Poll created successfully"
// @Failure 400 {object} utility.ErrorCode "4xx Errors"
// @Failure 500 {object} utility.ErrorCode "5xx Errors"
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

// AddVote godoc
// @Summary Vote on a poll
// @Description Cast a vote to a specific poll category by providing the poll ID in the route and the category ID in the request body.
// @Tags Polls
// @Accept json
// @Produce json
// @Param id path int true "Poll ID"
// @Param request body requests.AddVoteRequest true "Add Vote Request"
// @Success 200 {object} bool "Vote cast successful"
// @Failure 400 {object} utility.ErrorCode "4xx Errors"
// @Failure 500 {object} utility.ErrorCode "5xx Errors"
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

// DeletePoll godoc
// @Summary Delete a poll
// @Description Soft-Delete a poll by providing the poll ID in the route. You need to be the poll owner to delete it.
// @Tags Polls
// @Accept json
// @Produce json
// @Param id path int true "Poll ID"
// @Success 200 {object} bool "Poll deleted successfully"
// @Failure 400 {object} utility.ErrorCode "4xx Errors"
// @Failure 500 {object} utility.ErrorCode "5xx Errors"
// @Router /polls/{id} [delete]
// @Security BearerAuth
func (uc *PollController) DeletePoll(c *gin.Context) {

	pollIDString := c.Param("id")
	pollID, errParse := strconv.ParseUint(pollIDString, 10, 32)
	if errParse != nil {
		c.JSON(utility.RouteParameterCast.StatusCode, utility.RouteParameterCast)
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

	result, err := uc.DeletePollCommmand.DeletePoll(uint(pollID), user)

	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// EndPoll godoc
// @Summary End a poll
// @Description Mark the poll ended by providing the poll ID in the route. In order to end a poll, you need to be the poll owner.
// @Tags Polls
// @Accept json
// @Produce json
// @Param id path int true "Poll ID"
// @Success 200 {object} bool "Poll ended successfully"
// @Failure 400 {object} utility.ErrorCode "4xx Errors"
// @Failure 500 {object} utility.ErrorCode "5xx Errors"
// @Router /polls/{id}/end [patch]
// @Security BearerAuth
func (uc *PollController) EndPoll(c *gin.Context) {

	pollIDString := c.Param("id")
	pollID, errParse := strconv.ParseUint(pollIDString, 10, 32)
	if errParse != nil {
		c.JSON(utility.RouteParameterCast.StatusCode, utility.RouteParameterCast)
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

	result, err := uc.EndPollCommand.EndPoll(uint(pollID), user)

	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetPoll godoc
// @Summary Get a specific poll
// @Description Retrieve a specific poll by its ID.
// @Tags Polls
// @Accept json
// @Produce json
// @Param id path int true "Poll ID"
// @Success 200 {object} results.GetPollResult "Poll data"
// @Failure 400 {object} utility.ErrorCode "4xx Errors"
// @Failure 500 {object} utility.ErrorCode "5xx Errors"
// @Router /polls/{id} [get]
// @Security BearerAuth
func (uc *PollController) GetPoll(c *gin.Context) {

	pollIDString := c.Param("id")
	pollID, errParse := strconv.ParseUint(pollIDString, 10, 32)
	if errParse != nil {
		c.JSON(utility.RouteParameterCast.StatusCode, utility.RouteParameterCast)
		return
	}

	result, err := uc.GetPollQuery.GetPoll(uint(pollID))

	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetPolls godoc
// @Summary Get polls with pagination and optional filter
// @Description Retrieves a paginated list of polls and specifies whether to show only active polls. The filter parameter is used to search for polls by title or description.
// @Tags Polls
// @Accept json
// @Produce json
// @Param page query int false "Page number (default 1)"
// @Param page_size query int false "Items per page (default 10)"
// @Param filter query string false "Filter text (partial match against title or description)"
// @Param show_active_only query bool false "Show only active polls (default false)"
// @Success 200 {object} utility.PaginatedResponse[results.GetPollResult] "List of polls"
// @Failure 500 {object} utility.ErrorCode "Internal server error"
// @Router /polls [get]
func (uc *PollController) GetPolls(c *gin.Context) {

	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	filter := c.Query("filter")

	showActiveOnlyStr := c.Query("show_active_only")
	showActiveOnly, parseError := strconv.ParseBool(showActiveOnlyStr)
	if parseError != nil {
		c.JSON(utility.QueryParameterCast.StatusCode, utility.QueryParameterCast)
	}

	request := requests.GetPollsRequest{
		QueryParams: utility.QueryParams{
			Page:     page,
			PageSize: pageSize,
			Filter:   filter,
		},
		ShowActiveOnly: showActiveOnly,
	}

	result, err := uc.GetPollsQuery.GetPolls(&request)

	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetUserPolls godoc
// @Summary Get polls created by a specific user, with pagination/filter
// @Description Retrieves polls created by a specific user, with pagination and optional filter. The filter parameter is used to search for polls by title or description. The show_active_only parameter is used to show only active polls.
// @Tags Polls
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param page query int false "Page number (default 1)"
// @Param page_size query int false "Items per page (default 10)"
// @Param filter query string false "Filter text (partial match)"
// @Param show_active_only query bool false "Show only active polls (default false)"
// @Success 200 {object} utility.PaginatedResponse[results.GetPollResult] "List of user's polls"
// @Failure 400 {object} utility.ErrorCode "4xx Errors"
// @Failure 500 {object} utility.ErrorCode "5xx Errors"
// @Router /polls/users/{id} [get]
// @Security BearerAuth
func (uc *PollController) GetUserPolls(c *gin.Context) {

	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	filter := c.Query("filter")

	showActiveOnlyStr := c.Query("show_active_only")
	showActiveOnly, parseError := strconv.ParseBool(showActiveOnlyStr)
	if parseError != nil {
		c.JSON(utility.QueryParameterCast.StatusCode, utility.QueryParameterCast)
	}

	request := requests.GetPollsRequest{
		QueryParams: utility.QueryParams{
			Page:     page,
			PageSize: pageSize,
			Filter:   filter,
		},
		ShowActiveOnly: showActiveOnly,
	}

	userAny, ok := c.Get("user")
	if !ok {
		c.JSON(utility.Unauthorized.StatusCode, utility.Unauthorized)
	}

	user, ok := userAny.(*entities.User)
	if !ok {
		c.JSON(utility.InternalServerError.StatusCode, utility.InternalServerError)
	}

	result, err := uc.GetUserPollsQuery.GetPolls(user.ID, &request)

	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// UpdatePoll godoc
// @Summary Update a poll's details
// @Description Updates the specified poll's details, including title, expiration date, and categories.
// @Tags Polls
// @Accept json
// @Produce json
// @Param id path int true "Poll ID"
// @Param body body requests.UpdatePollRequest true "Poll update details"
// @Success 200 {object} bool "Poll updated successfully"
// @Failure 400 {object} utility.ErrorCode "4xx Errors"
// @Failure 500 {object} utility.ErrorCode "5xx Errors"
// @Router /polls/{id} [put]
// @Security BearerAuth
func (uc *PollController) UpdatePoll(c *gin.Context) {

	var request requests.UpdatePollRequest

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

	result, err := uc.UpdatePollCommand.UpdatePoll(user.ID, &request)

	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, result)
}
