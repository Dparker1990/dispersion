package gossip

import (
	"github.com/Dparker1990/dispersion/config"
	"testing"
)

func TestMerge(t *testing.T) {
	hash := make(map[string]Node)
	conf := config.Config{Bindip: "127.0.0.1", Bindport: "9292"}
	node := NewNode(conf)

	hash["127.0.0.1"] = *node
	node.Merge(hash)

	if _, ok := node.Peers["127.0.0.1"]; ok != true {
		t.Errorf("Expected hash to contain key, got: %v", node.Peers)
	}

	if node.Name() != "127.0.0.1:9292" {
		t.Errorf("Expected 127.0.0.1:9292, got: %v", node.Name())
	}
}
