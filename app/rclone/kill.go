package rclone

import (
	"os"

	"github.com/anotherhope/rcloud/app/internal"
)

// Kill running process and delete socket before exit
func Kill() {
	for _, process := range multiton {
		process.Command.Process.Kill()
	}
	os.Remove(internal.SocketPath)
}
