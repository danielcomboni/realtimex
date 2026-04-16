package realtimex

import "time"

type Event struct {
	Type      string      `json:"type"`
	Channel   string      `json:"channel,omitempty"`
	Payload   interface{} `json:"payload"`
	Timestamp time.Time   `json:"timestamp"`
}

func NewEvent(eventType string, payload interface{}) Event {
	return Event{
		Type:      eventType,
		Payload:   payload,
		Timestamp: time.Now(),
	}
}