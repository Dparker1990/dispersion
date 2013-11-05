package gossip

const (
	HEARTBEAT = 'h'
	SEED      = 's'
)

type Message struct {
	Length int32
	Type   byte
	Body   map[string]Node
}

func NewMessage(t int, body map[string]Node) Message {
	return Message{Type: HEARTBEAT, Body: body}
}
