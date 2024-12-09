package injection

import (
	"gin/api/authentication"
	"gin/application/repository"
	"gin/application/usecase/authentication/commands"
	"gin/infrastructure/database"
)

type AppContainer struct {
	AuthenticationController *authentication.AuthenticationController
}

// Set up all the dependencies & return controllers
func BuildContainer() *AppContainer {

	// Database connection
	database, _ := database.NewDatabase()
	dbContext := database.GetDBContext()

	UserRepository := repository.NewUserRepository(dbContext)
	RegisterCommand := commands.NewRegisterCommand(UserRepository)
	AuthenticationController := authentication.NewAuthenticationController(RegisterCommand)

	return &AppContainer{
		AuthenticationController: AuthenticationController,
	}
}
