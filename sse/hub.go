package sse

import (
	"encoding/json"
	"sync"

	"github.com/danielcomboni/realtimex"
)

type Hub struct {
	clients map[string]chan string
	mu      sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		clients: make(map[string]chan string),
	}
}

func (h *Hub) AddClient(clientID string, client interface{}) {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.clients[clientID] = client.(chan string)
}

func (h *Hub) RemoveClient(clientID string) {
	h.mu.Lock()
	defer h.mu.Unlock()

	delete(h.clients, clientID)
}

func (h *Hub) Broadcast(event realtimex.Event) error {
	data, _ := json.Marshal(event)

	h.mu.Lock()
	defer h.mu.Unlock()

	for _, ch := range h.clients {
		ch <- string(data)
	}

	return nil
}

func (h *Hub) Send(clientID string, event realtimex.Event) error {
	data, _ := json.Marshal(event)

	h.mu.Lock()
	defer h.mu.Unlock()

	if ch, ok := h.clients[clientID]; ok {
		ch <- string(data)
	}

	return nil
}
