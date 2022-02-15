package rclone

import (
	"time"

	"github.com/anotherhope/rcloud/app/config"
)

// Daemon contains the algo for sync
func Daemon(d *config.Directory) {
	var queue chan string = make(chan string)
	var lock bool = false

	go func() {
		for action := range queue {
			lock = true
			switch action {
			case "check":
				d.SetStatus("check")
				go func() {
					queue <- Check(d)
				}()
			case "sync":
				d.SetStatus("sync")
				go func() {
					queue <- Sync(d)
				}()
			case "idle":
				d.SetStatus("idle")
			}
			lock = false
		}
	}()

	if d.IsLocal(d.Source) && !d.IsLocal(d.Destination) {
		go runLocalChange(d, lock, queue)
	} else {
		go runRemoteChange(lock, queue)
	}

}

func runRemoteChange(lock bool, queue chan string) {
	for {
		if !lock {
			queue <- "check"
		}
		time.Sleep(5 * time.Second)
	}
}

func runLocalChange(d *config.Directory, lock bool, queue chan string) {
	for action := range d.CreateMirror(d.Source) {
		if !lock && d.SourceHasChange(action) {
			queue <- "sync"
		}
	}
}
