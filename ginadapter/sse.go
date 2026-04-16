package ginadapter

import (
	"fmt"

	"github.com/danielcomboni/realtimex/sse"
	"github.com/gin-gonic/gin"
)

func SSEHandler(hub *sse.Hub) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientID := c.Query("clientId")
		stream := make(chan string)

		hub.AddClient(clientID, stream)
		defer hub.RemoveClient(clientID)

		c.Writer.Header().Set("Content-Type", "text/event-stream")
		c.Writer.Header().Set("Cache-Control", "no-cache")
		c.Writer.Header().Set("Connection", "keep-alive")

		for msg := range stream {
			fmt.Fprintf(c.Writer, "data: %s\n\n", msg)
			c.Writer.Flush()
		}
	}
}
