package realtimex

import "sync"

type ChannelRegistry struct {
	channels map[string]map[string]bool
	mu       sync.RWMutex
}

func NewChannelRegistry() *ChannelRegistry {
	return &ChannelRegistry{
		channels: make(map[string]map[string]bool),
	}
}

func (r *ChannelRegistry) Subscribe(channel, clientID string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.channels[channel]; !ok {
		r.channels[channel] = make(map[string]bool)
	}

	r.channels[channel][clientID] = true
}

func (r *ChannelRegistry) Members(channel string) []string {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var ids []string

	for id := range r.channels[channel] {
		ids = append(ids, id)
	}

	return ids
}
