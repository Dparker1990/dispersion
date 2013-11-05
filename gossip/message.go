package gossip

const (
	HEARTBEAT = 'h'
)

type Message struct {
	Length int32
	Type   byte
	Body   map[string]Node
}
