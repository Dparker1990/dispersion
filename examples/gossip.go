package main

import (
	"flag"
	"fmt"
	"github.com/Dparker1990/dispersion/gossip"
	"os"
	"time"
)

func blather(node *gossip.Node) {
	for {
		node.Gossip()
		fmt.Println(fmt.Sprintf("Hash is now: %v", node.Peers))
		time.Sleep(1 * time.Second)
	}
}

func main() {
	var seed bool
	flag.BoolVar(&seed, "s", false, "Whether or not to act as a seed.")
	flag.Parse()

	s := os.Getenv("GOSSIP_SEED")
	config := gossip.Config{Bindip: "127.0.0.1", Bindport: "9292", Seeds: []string{s}}
	node := gossip.NewNode(config)

	if seed {
		go node.StartServer()
	} else {
		node.Register()
	}
	blather(node)
}
