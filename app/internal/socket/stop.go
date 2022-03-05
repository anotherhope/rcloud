package socket

import (
	"os"

	"github.com/anotherhope/rcloud/app/rclone"
)

func Stop() {
	rclone.Kill()
	os.Remove(SocketPath)
}
