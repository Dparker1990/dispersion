package gossip

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"io"
	"log"
	"net"
)

type MessageReader struct {
	conn net.Conn
}

func NewMessageReader(conn net.Conn) MessageReader {
	return MessageReader{conn: conn}
}

func (m *MessageReader) Read() (msg *Message, err error) {
	t := readByte(m.conn)
	l := readInt32(m.conn)
	log.Printf("Parsed message, type: %v, length: %v", t, l)

	b, err := decodeBody(readBody(m.conn, l))
	if err != nil {
		return
	}

	return &Message{Type: t, Length: l, Body: b}, nil
}

func decodeBody(encoded []byte) (body map[string]Node, err error) {
	buf := bytes.NewBuffer(encoded)
	decoder := gob.NewDecoder(buf)
	err = decoder.Decode(&body)

	return
}

func readByte(conn net.Conn) byte {
	buf := make([]byte, 1)
	_, err := io.ReadFull(conn, buf)
	if err != nil {
		log.Fatalf("Could not read byte due to: %v", err)
	}

	return buf[0]
}

func readInt32(conn net.Conn) int32 {
	buf := make([]byte, 4)
	_, err := io.ReadFull(conn, buf)
	if err != nil {
		log.Fatalf("Could not read int32 due to: %v", err)
	}
	return int32(binary.BigEndian.Uint32(buf))
}

func readBody(conn net.Conn, size int32) []byte {
	buf := make([]byte, size)
	_, err := io.ReadFull(conn, buf)
	if err != nil {
		log.Fatalf("Could not read message body due to: %v", err)
	}

	return buf
}
