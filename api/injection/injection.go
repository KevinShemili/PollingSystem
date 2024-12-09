package injection

import (
	"gin/api/authentication"
	"gin/application/repository"
	"gin/application/repository/contracts"
	"gin/application/usecase/authentication/commands"
	"gin/infrastructure/database"
)

type AppContainer struct {
	UserRepository contracts.IUserRepository // inject in auth-middleware

	AuthenticationController *authentication.AuthenticationController
}

// Set up all the dependencies & return controllers
func BuildContainer() *AppContainer {

	// Database connection
	database, _ := database.NewDatabase()
	dbContext := database.GetDBContext()

	UserRepository := repository.NewUserRepository(dbContext)
	RegisterCommand := commands.NewRegisterCommand(UserRepository)
	LoginCommand := commands.NewLoginCommand(UserRepository)
	AuthenticationController := authentication.NewAuthenticationController(RegisterCommand, LoginCommand)

	return &AppContainer{
		UserRepository: UserRepository,

		AuthenticationController: AuthenticationController,
	}
}
