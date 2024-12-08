// Backend Code:

package main

import (
	"gin/api/initializers"
	"gin/infrastructure/database"
	"gin/infrastructure/websocket"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	initializers.LoadEnvVariables()
	database.ConnectToDB()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// WebSocket handler
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
