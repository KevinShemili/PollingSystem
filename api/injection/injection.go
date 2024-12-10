package injection

import (
	"gin/api/authentication"
	"gin/application/repository"
	"gin/application/repository/contracts"
	"gin/application/usecase/authentication/commands"
	"gin/infrastructure/database"
)

type AppContainer struct {
	UnitOfWork contracts.IUnitOfWork // needed in auth-middleware

	AuthenticationController *authentication.AuthenticationController
}

// Set up all the dependencies & return controllers
func BuildContainer() *AppContainer {

	// Database connection
	database, _ := database.NewDatabase()
	dbContext := database.GetDBContext()

	// unit of work - repositories inside
	UnitOfWork := repository.NewUnitOfWork(dbContext)

	// handlers
	RegisterCommand := commands.NewRegisterCommand(UnitOfWork)
	LoginCommand := commands.NewLoginCommand(UnitOfWork)
	RefreshCommand := commands.NewRefreshCommand(UnitOfWork)
	LogOutCommand := commands.NewLogOutCommand(UnitOfWork)

	// controllers
	AuthenticationController := authentication.NewAuthenticationController(RegisterCommand, LoginCommand, RefreshCommand, LogOutCommand)

	return &AppContainer{
		UnitOfWork: UnitOfWork,

		AuthenticationController: AuthenticationController,
	}
}
