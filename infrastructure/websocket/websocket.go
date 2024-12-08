package websocket

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow any origin for simplicity
	},
}

var clients = make(map[*websocket.Conn]bool) // Connected WebSocket clients
var broadcast = make(chan string)            // Broadcast channel
var mutex = &sync.Mutex{}                    // Protect the clients map

// UpgradeConnection upgrades an HTTP connection to a WebSocket connection.
func UpgradeConnection(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}

	// Register the new client
	mutex.Lock()
	clients[conn] = true
	mutex.Unlock()

	// Set up read deadline and pong handler
	conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	conn.SetPongHandler(func(appData string) error {
		conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})

	fmt.Println("New WebSocket connection established")
	return conn, nil
}

// DeregisterClient removes a WebSocket connection from the clients list.
func DeregisterClient(conn *websocket.Conn) {
	mutex.Lock()
	delete(clients, conn)
	mutex.Unlock()
	err := conn.Close()
	if err != nil {
		fmt.Printf("Error closing WebSocket connection: %v\n", err)
	}
}

// BroadcastMessage sends a message to the broadcast channel.
func BroadcastMessage(msg string) {
	broadcast <- msg
}

// HandleBroadcast continuously listens for messages on the broadcast channel and sends them to all connected clients.
func HandleBroadcast() {
	for {
		msg := <-broadcast
		mutex.Lock()
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil {
				fmt.Printf("Error writing message: %v. Deregistering client.\n", err)
				client.Close()
				delete(clients, client)
			}
		}
		mutex.Unlock()
	}
}
