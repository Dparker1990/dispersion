package gossip

import (
	"bytes"
	"encoding/gob"
	"testing"
)

func TestDecodeBody(t *testing.T) {
	buf := new(bytes.Buffer)
	hash := make(map[string]Node)
	encoder := gob.NewEncoder(buf)

	hash["foo"] = Node{Health: "active"}
	if err := encoder.Encode(hash); err != nil {
		t.Errorf("Could not encode due to: %v", err)
	}

	newHash, err := decodeBody(buf.Bytes())
	if err != nil {
		t.Errorf("Could not decode body due to: %v", err)
	}

	node := newHash["foo"]
	if node.Health != "active" {
		t.Errorf("Decode did not happen properly, got: %v", newHash)
	}
}
