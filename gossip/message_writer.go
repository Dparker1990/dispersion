package gossip

import (
	"bytes"
	"encoding/gob"
	"net"
)

type MessageWriter struct {
	conn net.Conn
	buf  *bytes.Buffer
}

func NewMessageWriter(conn net.Conn) *MessageWriter {
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
