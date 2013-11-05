package gossip

import (
	"github.com/Dparker1990/dispersion/config"
	"log"
	"net"
)

const (
	ACTIVE     = iota
	SUSPICIOUS = iota
	DEAD       = iota
)

type Node struct {
	Health int
	Peers  map[string]Node
	Conf   config.Config
}

func NewNode() *Node {
	hash := make(map[string]Node)
	conf, err := config.Parse()
	if err != nil {
		log.Fatalf("Could not parse config due to: %v", err)
	}

	return &Node{Health: ACTIVE, Peers: hash, Conf: conf}
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
