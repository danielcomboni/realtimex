package realtimex

type Transport interface {
	Broadcast(event Event) error
	Send(clientID string, event Event) error
	AddClient(clientID string, client interface{})
	RemoveClient(clientID string)
}