package gossip

const (
	HEARTBEAT = 'h'
	SEED      = 's'
)

type Message struct {
	Type byte
	Body map[string]Node
}

func NewMessage(t byte, body map[string]Node) Message {
	return Message{Type: t, Body: body}
}
