package rcloud

import (
	"time"

	"github.com/anotherhope/rcloud/app/internal"
	"github.com/anotherhope/rcloud/app/rclone"
)

// Listen contains the algo for sync
func Listen(d *internal.Directory) {
	if d.GetChannel() == nil {
		var queue chan string = make(chan string)
		var lock bool = false

		go func() {
		runtime:
			for action := range queue {
				lock = true
				switch action {
				case "check":
					d.SetStatus("check")
					go func() {
						queue <- rclone.Check(d)
					}()
				case "sync":
					d.SetStatus("sync")
					go func() {
						queue <- rclone.Sync(d)
					}()
				case "exit":
					d.SetStatus("exit")
					d.SetChannel(nil)
					break runtime
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

		d.SetChannel(queue)
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

func runLocalChange(d *internal.Directory, lock bool, queue chan string) {
	d.SetStatus("idle")
	for action := range d.CreateMirror(d.Source) {
		if !lock && d.SourceHasChange(action) {
			queue <- "sync"
		}
	}
}
