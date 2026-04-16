package ws

import (
	"log"

	"github.com/gorilla/websocket"
)

func Listen(clientID string, conn *websocket.Conn, onMessage func([]byte)) {
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("client disconnected: %s", clientID)
			return
		}

		onMessage(message)
	}
}