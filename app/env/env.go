package env

import "os/user"

var SocketPath string
var User *user.User

func init() {
	User, _ = user.Current()
	SocketPath = User.HomeDir + "/.config/rcloud/daemon.sock"
}
