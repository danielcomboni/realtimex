package ginadapter

import (
	"encoding/json"
	"net/http"

	"github.com/danielcomboni/realtimex"
	"github.com/danielcomboni/realtimex/ws"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WSHandler(hub *ws.Hub) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientID := c.Query("clientId")

		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			c.AbortWithStatus(500)
			return
		}

		hub.AddClient(clientID, conn)
		defer hub.RemoveClient(clientID)

		ws.Listen(clientID, conn, func(message []byte) {
			var incoming map[string]interface{}
			_ = json.Unmarshal(message, &incoming)

			event := realtimex.NewEvent(
				"client.message.received",
				incoming,
			)

			_ = hub.Send(clientID, event)
		})
	}
}
