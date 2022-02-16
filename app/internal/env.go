package internal

import (
	"os"
	"os/user"
	"path"
)

// SocketPath is the path of socket file
var SocketPath string

// CachePath is the path of cache folder
var CachePath string

// User is the current running user
var User *user.User

func init() {
	User, _ = user.Current()
	SocketPath = User.HomeDir + "/.config/rcloud/daemon.sock"
	CachePath = User.HomeDir + "/.config/rcloud/cache"

	os.MkdirAll(path.Dir(SocketPath), 0700)
	os.MkdirAll(CachePath, 0700)
}
