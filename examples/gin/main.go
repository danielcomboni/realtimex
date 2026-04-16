package main

import (
	"time"

	"github.com/danielcomboni/realtimex"
	"github.com/danielcomboni/realtimex/ginadapter"
	"github.com/danielcomboni/realtimex/sse"
	"github.com/danielcomboni/realtimex/ws"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"

)

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://127.0.0.1:5500",
			"http://localhost:5500",
			"http://localhost:5173",
			"https://lainisha.local",
		},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	manager := realtimex.NewManager()

	wsHub := ws.NewHub()
	sseHub := sse.NewHub()

	manager.Register("chat", wsHub)
	manager.Register("stream", sseHub)

	r.GET("/api/ws", ginadapter.WSHandler(wsHub))
	r.GET("/api/stream", ginadapter.SSEHandler(sseHub))

	go func() {
		for {
			time.Sleep(2 * time.Second)

			event := realtimex.NewEvent(
				"heartbeat",
				map[string]string{
					"status": "alive",
				},
			)

			_ = manager.Broadcast("chat", event)
			_ = manager.Broadcast("stream", event)
		}
	}()

	r.Run(":7006")
}
