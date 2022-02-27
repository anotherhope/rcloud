package socket

import (
	"io/ioutil"
	"net"

	"github.com/anotherhope/rcloud/app/internal/socket/message"
)

// Dial implement connection and feature around socket communication
type Dial struct {
	conn net.Conn
}

// Close connection
func (d *Dial) Close() {
	d.conn.Close()
}

// Read response from server
func (d *Dial) Read() *message.Response {
	bytes, err := ioutil.ReadAll(d.conn)
	if err != nil {
		panic(err)
	}

	return &message.Response{
		Data: bytes,
	}
}

// Send a Message to server
func (d *Dial) Send(m *message.Message) {
	d.conn.Write(m.Request.ToBytes())
	m.Response = d.Read()
}

// Client create a socket client
func Client() *Dial {
	conn, err := net.Dial("unix", SocketPath)
	if err != nil {
		return nil
	}

	client := &Dial{
		conn: conn,
	}

	return client
}
