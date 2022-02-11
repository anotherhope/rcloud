package rclone

import (
	"github.com/anotherhope/rcloud/app/config"
)

func Daemon(d *config.Directory) {
	for {
		if Check(d) {
			Sync(d)
		}
	}
}
