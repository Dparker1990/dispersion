package gossip

import (
	"fmt"
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
	peer, err := n.randomPeer()
	if err != nil {
		return
	}

	conn, err := net.Dial("tcp", peer)
	defer conn.Close()
	if err != nil {
		log.Printf("Could not contact %v", peer)
		return
	}

	msg := NewMessage(HEARTBEAT, n.Peers)
	messageWriter := NewMessageWriter(conn)
	if _, err = messageWriter.Write(msg); err != nil {
		log.Printf("Could not send message due to: %v", err)
	}
}

func (n *Node) Register() {
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

	n.Peers[seed] = Node{Health: ACTIVE}
}

func (n *Node) StartServer() (err error) {
	ln, err := net.Listen("tcp", n.Conf.Bindport)
	if err != nil {
		panic(fmt.Sprintf("Could not listen due to: %v", err))
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
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

func (n Node) AllPeerKeys() (keys []string) {
	for k, _ := range n.Peers {
		keys = append(keys, k)
	}

	return
}

func (n *Node) randomPeer() (peer string, err error) {
	keys := n.AllPeerKeys()
	max := len(keys)
	if max == 0 {
		err = fmt.Errorf("No peers.")
		return
	}

	idx := rand.Intn(max)
	if idx == max {
		idx--
	}

	peer = keys[idx]
	return
}
