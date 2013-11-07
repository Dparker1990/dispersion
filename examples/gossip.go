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
		if err := node.Gossip(); err != nil {
			fmt.Print(err)
		}
		fmt.Printf("Hash is now: %v\n", node.Peers)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	var seed bool
	flag.BoolVar(&seed, "s", false, "Whether or not to act as a seed.")
	flag.Parse()

	node := gossip.NewNode()
	s := os.Getenv("GOSSIP_SEED")
	config := gossip.Config{Bindip: "72.54.29.74", Bindport: "9292", Seeds: []string{s}}
	node.SetConfig(config)

	if !seed {
		node.Register()
	}
	go node.StartServer()
	blather(node)
}
