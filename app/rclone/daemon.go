package rclone

import (
	"time"

	"github.com/anotherhope/rcloud/app/config"
)

// Daemon contains the algo for sync
func Daemon(d *config.Directory) {
	for {
		if Check(d) {
			Sync(d)
		}

		time.Sleep(1 * time.Second)
	}
}
