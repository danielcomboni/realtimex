package realtimex

type Manager struct {
	transports map[string]Transport
}

func NewManager() *Manager {
	return &Manager{
		transports: make(map[string]Transport),
	}
}

func (m *Manager) Register(name string, transport Transport) {
	m.transports[name] = transport
}

func (m *Manager) Broadcast(name string, event Event) error {
	if t, ok := m.transports[name]; ok {
		return t.Broadcast(event)
	}
	return nil
}

func (m *Manager) BroadcastTo(name string, clientIDs []string, event Event) error {
	if t, ok := m.transports[name]; ok {
		for _, id := range clientIDs {
			_ = t.Send(id, event)
		}
	}
	return nil
}
