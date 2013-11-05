package gossip

import (
	"log"
	"net"
)

const (
	HEARTBEAT = 'h'
)

type Node struct {
	Health string
	Peers  map[string]Node
}

func NewNode() *Node {
	hash := make(map[string]Node)
	return &Node{Health: "active", Peers: hash}
}

func (n *Node) HandleConnection(conn net.Conn) {
	defer conn.Close()

	msgReader := MessageReader{conn: conn}
	msg, err := msgReader.Read()
	if err != nil {
		log.Printf("Could not read message due to: %v", err)
	}

	n.Merge(msg.Body)
}

func (n *Node) Merge(hash map[string]Node) {
	for key, value := range hash {
		if _, ok := n.Peers[key]; ok != true {
			n.Peers[key] = value
		}
	}
}
