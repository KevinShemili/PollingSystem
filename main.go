package main

import (
	"gin/api/authentication"
	"gin/api/initializers"
	"gin/api/injection"
	"gin/infrastructure/websocket"
	"net/http"

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

	authentication.AuthenticationRoutes(r, container.AuthenticationController, container.UnitOfWork)

	// -------------------------------------------------------------------------------------

	// WEBSOCKETS
	// Keep the connection open without reading messages
	r.GET("/ws", func(c *gin.Context) {
		conn, err := websocket.UpgradeConnection(c.Writer, c.Request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "WebSocket upgrade failed"})
			return
		}
		defer websocket.DeregisterClient(conn)

		// Block indefinitely
		select {}
	})

	// Simple controller to broadcast junk
	r.GET("/broadcast", func(c *gin.Context) {
		websocket.BroadcastMessage("This is junk data!") // Send junk data to the broadcast channel
		c.JSON(200, gin.H{"message": "Broadcast sent"})
	})

	// Start broadcaster in a goroutine
	go websocket.HandleBroadcast()

	r.Run()
}
