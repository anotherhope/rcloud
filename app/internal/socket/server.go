package socket

import (
	"bufio"
	"net"

	"github.com/anotherhope/rcloud/app/internal/socket/action"
)

// Client create a socket server
func Server() error {
	ln, err := net.Listen("unix", SocketPath)

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
