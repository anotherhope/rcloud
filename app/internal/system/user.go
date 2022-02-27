package system

import (
	"os/user"
)

// User is the current running user
var User *user.User

func init() {
	User, _ = user.Current()
}
