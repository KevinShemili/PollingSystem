package authentication

import (
	"github.com/gin-gonic/gin"
)

func AuthenticationRoutes(r *gin.Engine, controller *AuthenticationController) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", controller.Register)
		auth.GET("/login", controller.Login)
		auth.POST("/logout", controller.LogOut)
	}
}
