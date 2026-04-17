package main

import (
	"time"

	"github.com/danielcomboni/realtimex"
	"github.com/danielcomboni/realtimex/ginadapter"
	"github.com/danielcomboni/realtimex/sse"
	"github.com/danielcomboni/realtimex/ws"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	manager := realtimex.NewManager()

	wsHub := ws.NewHub()
	sseHub := sse.NewHub()

	// manager.Register("chat", wsHub)
	manager.Register("stream", sseHub)

	// manager.Register("orders", wsHub)
	manager.Register("dashboard", sseHub)

	// r.GET("/api/ws", ginadapter.WSHandler(wsHub))
	r.GET("/api/stream", ginadapter.SSEHandler(sseHub))

	r.GET("/api/ws/orders", ginadapter.WSHandler(wsHub))
	r.GET("/api/stream/dashboard", ginadapter.SSEHandler(sseHub))
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

	event := realtimex.NewEvent(
		"heartbeat2",
		map[string]string{
			"status": "alive2",
		},
	)

	scheduler := realtimex.NewScheduler()

	scheduler.Every(2*time.Second, func() {
		manager.Broadcast("chat", event)
		manager.Broadcast("stream", event)
	})

	// Simulate new orders coming in
	go func() {
		orderNo := 1000

		for {
			time.Sleep(5 * time.Second)

			orderNo++

			event := realtimex.NewEvent(
				"order.created",
				map[string]interface{}{
					"orderId": orderNo,
					"table":   "A1",
					"amount":  45.50,
					"status":  "pending",
				},
			)

			// _ = manager.Broadcast("orders", event)
			_ = manager.Broadcast("dashboard", event)
		}
	}()

	r.Run(":7006")
}
