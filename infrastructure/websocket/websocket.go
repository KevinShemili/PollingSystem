package websocket

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// this is used to upgrade HTTP -> WebSocket
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan string)
var mutex = &sync.Mutex{}

func UpgradeConnection(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {

	// websocket handshake
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}

	// Register the new client
	mutex.Lock()
	clients[conn] = true
	mutex.Unlock()

	fmt.Println("New websocket client.")
	return conn, nil
}

// place message in broadcast channel -> broadcast to all clients
func BroadcastMessage(msg string) {
	broadcast <- msg
}

// continuously listens for new mssg on channel -> sends them to connected clients
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
