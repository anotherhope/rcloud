package socket

import (
	"errors"
	"os"
	"path"

	"github.com/anotherhope/rcloud/app/internal/system"
)

// SocketPath is the path of socket file
var SocketPath string

func init() {
	SocketPath = system.User.HomeDir + "/.config/rcloud/daemon.sock"

	if _, err := os.Stat(SocketPath); errors.Is(err, os.ErrNotExist) {
		os.MkdirAll(path.Dir(SocketPath), 0700)
	}
}
