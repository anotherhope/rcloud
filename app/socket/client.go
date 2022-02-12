package socket

import (
	"io/ioutil"
	"net"

	"github.com/anotherhope/rcloud/app/env"
	"github.com/anotherhope/rcloud/app/message"
)

type Dial struct {
	conn net.Conn
}

func (d *Dial) Close() {
	d.conn.Close()
}

func (d *Dial) Read() *message.Response {
	bytes, err := ioutil.ReadAll(d.conn)
	if err != nil {
		panic(err)
	}

	return &message.Response{
		Data: bytes,
	}
}

func (d *Dial) Send(m *message.Message) {
	d.conn.Write(m.Request.ToBytes())
	m.Response = d.Read()
}

func Client() *Dial {
	conn, err := net.Dial("unix", env.SocketPath)
	if err != nil {
		return nil
	}

	client := &Dial{
		conn: conn,
	}

	return client
}
