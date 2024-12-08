package injection

import (
	"gin/api/controllers"
	"gin/application/repository"
	"gin/application/usecase/authentication/commands"
	"gin/infrastructure/database"
)

type AppContainer struct {
	AuthenticationController *controllers.AuthenticationController
}

// Set up all the dependencies & return controllers
func BuildContainer() *AppContainer {

	// Database connection
	database := database.ConnectToDB()
	gormInstace := database.GetDBContext()

	UserRepository := repository.NewUserRepository(gormInstace)
	RegisterUserCommand := commands.NewRegisterUserCommand(UserRepository)
	AuthenticationController := controllers.NewAuthenticationController(RegisterUserCommand)

	return &AppContainer{
		AuthenticationController: AuthenticationController,
	}
}
