package gossip

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"net"
)

type MessageWriter struct {
	conn net.Conn
	buf  *bytes.Buffer
	msg  Message
}

func NewMessageWriter(conn net.Conn) *MessageWriter {
	buf := new(bytes.Buffer)
	return &MessageWriter{conn: conn, buf: buf}
}

func (m *MessageWriter) Write(msg Message) (size int, err error) {
	writer := bufio.NewWriter(m.conn)
	m.msg = msg

	if err = writer.WriteByte(msg.Type); err != nil {
		return
	}

	size, err = m.sendBody()
	return
}

func (m *MessageWriter) sendBody() (size int, err error) {
	encoder := gob.NewEncoder(m.buf)

	if err = encoder.Encode(m.msg.Body); err != nil {
		return
	}

	length := m.buf.Len()
	if err = m.writeInt32(length); err != nil {
		return
	}

	size, err = m.conn.Write(m.buf.Bytes())
	return
}

func (m *MessageWriter) writeInt32(n int) error {
	err := binary.Write(m.conn, binary.BigEndian, int32(n))
	return err
}
