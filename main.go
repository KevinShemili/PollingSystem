package main

import (
	"fmt"
	"gin/api/initializers"
	"gin/api/injection"
	"gin/api/routes"
	"gin/application/usecase/poll/commands"
	"gin/infrastructure/websocket"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	_ "gin/docs" // Must import for Swagger to render

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Your API Title
// @version 1.0
// @description Your API description
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {

	initializers.LoadEnvironmentVariabes()

	container := injection.BuildContainer()

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.AuthenticationRoutes(r, container.AuthenticationController, container.UnitOfWork)
	routes.PollRoutes(r, container.PollController, container.UnitOfWork)

	// -------------------------------------------------------------------------------------

	// WEBSOCKET
	r.GET("/ws", func(c *gin.Context) {

		_, err := websocket.UpgradeConnection(c.Writer, c.Request)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "WebSocket upgrade failed"})
			return
		}

		// Block indefinitely
		select {}
	})

	// broadcaster in a goroutine
	go websocket.HandleBroadcast()

	// routine for poll expiration
	go func() {
		// every 1 minute - check for expiries
		ticker := time.NewTicker(1 * time.Minute)

		for range ticker.C {
			if err := commands.EndExpiredPolls(container.UnitOfWork); err != nil {
				fmt.Printf("Poll Expiry Error: %v", err)
			}
		}
	}()

	r.Run()
}
