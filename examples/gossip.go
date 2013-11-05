package main

import (
	"flag"
	"github.com/Dparker1990/dispersion/gossip"
	"os"
	"time"
)

func blather(node *gossip.Node) {
	node.Gossip()
	time.Sleep(1 * time.Second)
}

func main() {
	seed := flag.Bool("-s", false, "Whether or not to act as a seed.")
	flag.Parse()

	s := os.Getenv("GOSSIP_SEED")
	config := gossip.Config{Bindip: "127.0.0.1", Bindport: "9292", Seeds: []string{s}}
	node := gossip.NewNode(config)

	if !*seed {
		node.Register()
	}
	go blather(node)
	go node.StartServer()
}
