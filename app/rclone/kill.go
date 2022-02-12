package rclone

import (
	"os"

	"github.com/anotherhope/rcloud/app/env"
)

func Kill() {
	for _, process := range multiton {
		process.Command.Process.Kill()
	}
	os.Remove(env.SocketPath)
}
