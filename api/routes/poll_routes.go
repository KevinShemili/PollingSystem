/*
Poll Routes
Auth Middleware used to protect some resources
*/

package routes

import (
	"gin/api/controllers"
	"gin/api/middleware"
	"gin/application/repository/contracts"

	"github.com/gin-gonic/gin"
)

func PollRoutes(r *gin.Engine, controller *controllers.PollController, UnitOfWork contracts.IUnitOfWork) {

	poll := r.Group("/polls")
	{
		poll.POST("", middleware.AuthenticationMiddleware(UnitOfWork), controller.CreatePoll)
		poll.POST("/:id/vote", middleware.AuthenticationMiddleware(UnitOfWork), controller.AddVote)
		poll.DELETE("/:id", middleware.AuthenticationMiddleware(UnitOfWork), controller.DeletePoll)
		poll.PATCH("/:id/end", middleware.AuthenticationMiddleware(UnitOfWork), controller.EndPoll)
		poll.GET("", controller.GetPolls)
		poll.GET("/:id", middleware.AuthenticationMiddleware(UnitOfWork), controller.GetPoll)
		poll.GET("/users/:id", middleware.AuthenticationMiddleware(UnitOfWork), controller.GetUserPolls)
		poll.PUT("/:id", middleware.AuthenticationMiddleware(UnitOfWork), controller.UpdatePoll)
	}
}
