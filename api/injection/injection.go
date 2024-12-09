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
	database := database.ConnectToDB()
	gormInstace := database.GetDBContext()

	UserRepository := repository.NewUserRepository(gormInstace)
	RegisterUserCommand := commands.NewRegisterUserCommand(UserRepository)
	AuthenticationController := authentication.NewAuthenticationController(RegisterUserCommand)

	return &AppContainer{
		AuthenticationController: AuthenticationController,
	}
}
