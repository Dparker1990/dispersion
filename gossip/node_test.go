package gossip

import (
	"testing"
)

func TestNodeName(t *testing.T) {
	conf := Config{Bindip: "127.0.0.1", Bindport: "9292"}
	node := NewNode()
	node.SetConfig(conf)

	if node.Name() != "127.0.0.1:9292" {
		t.Errorf("Expected name as 127.0.0.1:9292, got: %v", node.Name())
	}
}

func TestMerge(t *testing.T) {
	hash := make(map[string]Node)
	conf := Config{Bindip: "127.0.0.1", Bindport: "9292"}
	node := NewNode()
	node.SetConfig(conf)

	hash["127.0.0.1"] = *node
	node.Merge(hash)

	if _, ok := node.Peers["127.0.0.1"]; ok != true {
		t.Errorf("Expected hash to contain key, got: %v", node.Peers)
	}

	if node.Name() != "127.0.0.1:9292" {
		t.Errorf("Expected 127.0.0.1:9292, got: %v", node.Name())
	}
}

func TestAllPeerKeys(t *testing.T) {
	hash := make(map[string]Node)
	conf := Config{Bindip: "127.0.0.1", Bindport: "9292"}
	node1 := NewNode()
	node2 := NewNode()
	node1.SetConfig(conf)
	node2.SetConfig(conf)

	hash["127.0.0.1"] = *node1
	hash["192.168.1.1"] = *node2
	node1.Merge(hash)

	allKeys := node1.AllPeerKeys()
	if len(allKeys) != 2 {
		t.Errorf("Keys did not come out as expected, got: %v", allKeys)
	}
}
