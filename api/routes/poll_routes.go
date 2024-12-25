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
		auth.POST("/:id/vote", middleware.AuthenticationMiddleware(UnitOfWork), controller.AddVote)
		auth.DELETE("/:id", middleware.AuthenticationMiddleware(UnitOfWork), controller.DeletePoll)
		auth.PUT("/:id/end", middleware.AuthenticationMiddleware(UnitOfWork), controller.EndPoll)
	}
}
