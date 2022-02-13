package env

import (
	"os"
	"os/user"
	"path"
)

var SocketPath string
var CachePath string
var User *user.User

func init() {
	User, _ = user.Current()
	SocketPath = User.HomeDir + "/.config/rcloud/daemon.sock"
	CachePath = User.HomeDir + "/.config/rcloud/cache"

	os.MkdirAll(path.Dir(SocketPath), 0700)
	os.MkdirAll(CachePath, 0700)
}
