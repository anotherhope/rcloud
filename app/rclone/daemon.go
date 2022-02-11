package rclone

import (
	"time"

	"github.com/anotherhope/rcloud/app/config"
)

func Daemon(d *config.Directory) {
	for {
		if Check(d) {
			Sync(d)
		}

		time.Sleep(1 * time.Second)
	}
}
