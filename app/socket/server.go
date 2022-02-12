package socket

import (
	"bufio"
	"net"

	"github.com/anotherhope/rcloud/app/env"
	"github.com/anotherhope/rcloud/app/socket/action"
)

func Server() error {
	ln, err := net.Listen("unix", env.SocketPath)

	if err != nil {
		return err
	}

	for {
		conn, _ := ln.Accept()
		message, _ := bufio.NewReader(conn).ReadString('\n')
		conn.Write(action.Do(message))
		conn.Close()
	}
}
