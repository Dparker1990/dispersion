package gossip

import (
	"log"
	"math/rand"
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
	Conf   Config
}

func NewNode(conf Config) *Node {
	hash := make(map[string]Node)

	return &Node{Health: ACTIVE, Peers: hash, Conf: conf}
}

func (n *Node) Gossip() {
	peer := n.randomPeer()

	conn, err := net.Dial("tcp", peer)
	defer conn.Close()

	if err != nil {
		log.Printf("Could not contact %v", peer)
	}

	msg := NewMessage(HEARTBEAT, n.Peers)
	messageWriter := NewMessageWriter(conn)
	if _, err = messageWriter.Write(msg); err != nil {
		log.Printf("Could not send message due to: %v", err)
	}
}

func (n Node) Register() {
	seed := n.Conf.Seeds[0]
	conn, err := net.Dial("tcp", seed)
	if err != nil {
		log.Fatalf("Could not connect to seed: %v", err)
	}
	defer conn.Close()

	msg := Message{Type: SEED}
	messageWriter := NewMessageWriter(conn)
	if _, err = messageWriter.Write(msg); err != nil {
		log.Printf("Could not send seed message due to: %v", err)
	}
}

func (n *Node) StartServer() (err error) {
	ln, err := net.Listen("tcp", n.Conf.Bindip)
	if err != nil {
		return
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			break
		}

		go n.HandleConnection(conn)
	}

	return
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

func (n Node) Name() string {
	return n.Conf.Bindip + ":" + n.Conf.Bindport
}

func (n *Node) Merge(hash map[string]Node) {
	for key, value := range hash {
		if _, ok := n.Peers[key]; ok != true {
			n.Peers[key] = value
		}
	}
}

func (n *Node) SetHealth(health int) {
	n.Health = health
}

func (n Node) randomPeer() string {
	keys := n.allPeerKeys()
	max := len(keys)
	idx := rand.Intn(max)
	if idx == max {
		idx--
	}

	return keys[idx]
}

func (n Node) allPeerKeys() (keys []string) {
	for k, _ := range n.Peers {
		keys = append(keys, k)
	}

	return
}
