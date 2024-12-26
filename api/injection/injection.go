package injection

import (
	"gin/api/controllers"
	"gin/application/repository"
	"gin/application/repository/contracts"
	authCommands "gin/application/usecase/authentication/commands"
	pollCommands "gin/application/usecase/poll/commands"
	pollQueries "gin/application/usecase/poll/queries"
	"gin/infrastructure/database"
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

	// handlers
	RegisterCommand := authCommands.NewRegisterCommand(UnitOfWork)
	LoginCommand := authCommands.NewLoginCommand(UnitOfWork)
	RefreshCommand := authCommands.NewRefreshCommand(UnitOfWork)
	LogOutCommand := authCommands.NewLogOutCommand(UnitOfWork)
	CreatePollCommand := pollCommands.NewCreatePollCommand(UnitOfWork)
	AddVoteCommand := pollCommands.NewAddVoteCommand(UnitOfWork)
	DeletePollCommand := pollCommands.NewDeletePollCommand(UnitOfWork)
	EndPollCommand := pollCommands.NewEndPollCommand(UnitOfWork)
	UpdatePollCommand := pollCommands.NewUpdatePollCommand(UnitOfWork)
	GetPollQuery := pollQueries.NewGetPollQuery(UnitOfWork)
	GetPollsQuery := pollQueries.NewGetPollsQuery(UnitOfWork)
	GetUserPollsQuery := pollQueries.NewGetUserPollsQuery(UnitOfWork)

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
