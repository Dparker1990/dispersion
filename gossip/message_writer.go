package gossip

import (
	"encoding/gob"
	"io"
)

type MessageWriter struct {
	conn io.Writer
}

func NewMessageWriter(conn io.Writer) *MessageWriter {
	return &MessageWriter{conn: conn}
}

func (m *MessageWriter) Write(msg Message) (size int, err error) {
	encoder := gob.NewEncoder(m.conn)
	if err = encoder.Encode(msg); err != nil {
		return
	}
	return
}
