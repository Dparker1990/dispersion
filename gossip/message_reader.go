package gossip

import (
	"encoding/gob"
	"io"
)

type MessageReader struct {
	conn io.Reader
}

func NewMessageReader(conn io.Reader) MessageReader {
	return MessageReader{conn: conn}
}

func (m *MessageReader) Read() (msg *Message, err error) {
	msg = &Message{}
	decoder := gob.NewDecoder(m.conn)
	err = decoder.Decode(msg)

	return
}
