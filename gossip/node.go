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
	Conn   net.Conn
	config Config
}

func NewNode() *Node {
	hash := make(map[string]Node)

	return &Node{Health: ACTIVE, Peers: hash}
}

func (n *Node) Config() Config {
	return n.config
}

func (n *Node) SetConfig(c Config) {
	n.config = c
}

func (n *Node) Gossip() (err error) {
	peer, err := n.randomPeer()
	if err != nil {
		return
	}

	conn := peer.Conn
	if err != nil {
		log.Printf("Could not contact %v", peer)
		return
	}

	msg := NewMessage(HEARTBEAT, n.Peers)
	messageWriter := NewMessageWriter(conn)
	if _, err = messageWriter.Write(msg); err != nil {
		log.Printf("Could not send message due to: %v", err)
	}

	return
}

func (n *Node) Register() {
	seed := n.Config().Seeds[0]
	conn, err := net.Dial("tcp", seed+":9292")
	if err != nil {
		log.Fatalf("Could not connect to seed: %v", err)
	}

	msg := Message{Type: SEED}
	messageWriter := NewMessageWriter(conn)
	if _, err = messageWriter.Write(msg); err != nil {
		log.Printf("Could not send seed message due to: %v", err)
	}

	n.Peers[seed] = Node{Health: ACTIVE}
}

func (n *Node) StartServer() (err error) {
	port := n.Config().Bindport
	log.Printf("Starting server, port: %v", port)
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(fmt.Sprintf("Could not listen due to: %v", err))
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}

		log.Printf("Received connection from: %v", conn.RemoteAddr())

		go n.HandleConnection(conn)
	}

	return
}

func (n *Node) HandleConnection(conn net.Conn) {
	msgReader := MessageReader{conn: conn}
	if err := n.addNewNode(conn); err != nil {
		log.Print(err)
	}

	for {
		msg, err := msgReader.Read()
		if err != nil {
			log.Printf("Could not read message due to: %v\n", err)
			return
		}

		n.Merge(msg.Body)
	}
}

func (n Node) Name() string {
	conf := n.Config()
	return net.JoinHostPort(conf.Bindip, conf.Bindport)
}

func (n *Node) Merge(hash map[string]Node) {
	for key, value := range hash {
		if _, ok := n.Peers[key]; !ok {
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

func (n *Node) addNewNode(conn net.Conn) (err error) {
	key, _, err := net.SplitHostPort(conn.RemoteAddr().String())
	if err != nil {
		log.Printf("Could not parse remote address due to: %v", err)
		return
	}

	if _, ok := n.Peers[key]; !ok {
		log.Printf("Adding node: %v", key)
		node := *NewNode()
		node.Conn = conn
		n.Peers[key] = node
	}

	return
}

func (n *Node) randomPeer() (peer Node, err error) {
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

	key := keys[idx]
	peer = n.Peers[key]
	return
}
