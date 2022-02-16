package socket

import (
	"bufio"
	"net"

	"github.com/anotherhope/rcloud/app/internal"
	"github.com/anotherhope/rcloud/app/socket/action"
)

// Client create a socket server
func Server() error {
	ln, err := net.Listen("unix", internal.SocketPath)

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
