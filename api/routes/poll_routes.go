package routes

import (
	"gin/api/controllers"
	"gin/api/middleware"
	"gin/application/repository/contracts"

	"github.com/gin-gonic/gin"
)

func PollRoutes(r *gin.Engine, controller *controllers.PollController, UnitOfWork contracts.IUnitOfWork) {

	auth := r.Group("/polls")
	{
		auth.POST("", middleware.AuthenticationMiddleware(UnitOfWork), controller.CreatePoll)
	}
}
