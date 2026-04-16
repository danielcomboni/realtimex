package ws

import (
	"sync"

	"github.com/danielcomboni/realtimex"
	"github.com/gorilla/websocket"
)

type Hub struct {
	clients map[string]*websocket.Conn
	mu      sync.RWMutex
}

func NewHub() *Hub {
	return &Hub{
		clients: make(map[string]*websocket.Conn),
	}
}

func (h *Hub) AddClient(clientID string, client interface{}) {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.clients[clientID] = client.(*websocket.Conn)
}

func (h *Hub) RemoveClient(clientID string) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if conn, ok := h.clients[clientID]; ok {
		_ = conn.Close()
		delete(h.clients, clientID)
	}
}

func (h *Hub) Broadcast(event realtimex.Event) error {
	h.mu.RLock()
	defer h.mu.RUnlock()

	for _, conn := range h.clients {
		if err := conn.WriteJSON(event); err != nil {
			_ = conn.Close()
		}
	}

	return nil
}

func (h *Hub) Send(clientID string, event realtimex.Event) error {
	h.mu.RLock()
	defer h.mu.RUnlock()

	if conn, ok := h.clients[clientID]; ok {
		return conn.WriteJSON(event)
	}

	return nil
}