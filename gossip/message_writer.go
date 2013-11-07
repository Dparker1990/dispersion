package gossip

import (
	"bytes"
	"encoding/gob"
	"io"
)

type MessageWriter struct {
	conn io.Writer
	buf  *bytes.Buffer
}

func NewMessageWriter(conn io.Writer) *MessageWriter {
	buf := new(bytes.Buffer)
	return &MessageWriter{conn: conn, buf: buf}
}

func (m *MessageWriter) Write(msg Message) (size int, err error) {
	encoder := gob.NewEncoder(m.conn)
	if err = encoder.Encode(msg); err != nil {
		return
	}
	return
}
