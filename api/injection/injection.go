package injection

import (
	"gin/api/controllers"
	"gin/application/repository"
	"gin/application/repository/contracts"
	authCommands "gin/application/usecase/authentication/commands"
	pollCommands "gin/application/usecase/poll/commands"
	pollQueries "gin/application/usecase/poll/queries"
	"gin/infrastructure/database"

	"github.com/go-playground/validator/v10"
)

type AppContainer struct {
	UnitOfWork contracts.IUnitOfWork // needed in auth-middleware

	AuthenticationController *controllers.AuthenticationController
	PollController           *controllers.PollController
}

// Set up all the dependencies & return controllers
func BuildContainer() *AppContainer {

	// Database connection
	database, _ := database.NewDatabase()
	dbContext := database.GetDBContext()

	// unit of work - repositories inside
	UnitOfWork := repository.NewUnitOfWork(dbContext)

	validate := validator.New()

	// handlers
	RegisterCommand := authCommands.NewRegisterCommand(UnitOfWork, validate)
	LoginCommand := authCommands.NewLoginCommand(UnitOfWork, validate)
	RefreshCommand := authCommands.NewRefreshCommand(UnitOfWork, validate)
	LogOutCommand := authCommands.NewLogOutCommand(UnitOfWork, validate)
	CreatePollCommand := pollCommands.NewCreatePollCommand(UnitOfWork, validate)
	AddVoteCommand := pollCommands.NewAddVoteCommand(UnitOfWork, validate)
	DeletePollCommand := pollCommands.NewDeletePollCommand(UnitOfWork)
	EndPollCommand := pollCommands.NewEndPollCommand(UnitOfWork)
	UpdatePollCommand := pollCommands.NewUpdatePollCommand(UnitOfWork, validate)
	GetPollQuery := pollQueries.NewGetPollQuery(UnitOfWork)
	GetPollsQuery := pollQueries.NewGetPollsQuery(UnitOfWork, validate)
	GetUserPollsQuery := pollQueries.NewGetUserPollsQuery(UnitOfWork, validate)

	// controllers
	AuthenticationController := controllers.NewAuthenticationController(RegisterCommand, LoginCommand, RefreshCommand, LogOutCommand)
	PollController := controllers.NewPollController(CreatePollCommand, AddVoteCommand, DeletePollCommand, EndPollCommand, GetPollQuery,
		GetPollsQuery, GetUserPollsQuery, UpdatePollCommand)

	return &AppContainer{
		UnitOfWork: UnitOfWork,

		AuthenticationController: AuthenticationController,
		PollController:           PollController,
	}
}
