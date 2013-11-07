package gossip

import (
	"bytes"
	"encoding/gob"
	"testing"
)

func TestWrite(t *testing.T) {
	buff := new(bytes.Buffer)
	writer := NewMessageWriter(buff)
	body := make(map[string]Node)
	body["foo"] = Node{Health: ACTIVE}
	msg := NewMessage(HEARTBEAT, body)
	decoder := gob.NewDecoder(buff)

	_, err := writer.Write(msg)
	if err != nil {
		t.Errorf("Expected write not to fail. %v", err)
	}

	var decodedMsg Message
	err = decoder.Decode(&decodedMsg)
	if err != nil {
		t.Errorf("Could not decode: %v", err)
	}

	if decodedMsg.Type != HEARTBEAT {
		t.Errorf("Message was not decoded properly, got: %v", decodedMsg)
	}
}
