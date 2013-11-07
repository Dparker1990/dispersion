package gossip

import (
	"bytes"
	"encoding/gob"
	"testing"
)

func TestDecodeBody(t *testing.T) {
	buf := new(bytes.Buffer)
	body := make(map[string]Node)
	sendMsg := &Message{Type: ACTIVE, Body: body}
	encoder := gob.NewEncoder(buf)
	reader := NewMessageReader(buf)

	body["foo"] = Node{Health: ACTIVE}
	if err := encoder.Encode(sendMsg); err != nil {
		t.Errorf("Could not encode due to: %v", err)
	}

	msg, err := reader.Read()
	if err != nil {
		t.Errorf("Could not decode body due to: %v", err)
	}

	node := msg.Body["foo"]
	if node.Health != ACTIVE {
		t.Errorf("Decode did not happen properly, got: %v", msg.Body)
	}
}
