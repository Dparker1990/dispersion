package gossip

import (
	"testing"
)

func TestMerge(t *testing.T) {
	hash := make(map[string]Node)
	node := NewNode()

	hash["127.0.0.1"] = *node
	node.Merge(hash)

	if _, ok := node.Peers["127.0.0.1"]; ok != true {
		t.Errorf("Expected hash to contain key, got: %v", node.Peers)
	}
}