package rcloud

import (
	"os"
	"time"

	"github.com/anotherhope/rcloud/app/internal"
	"github.com/anotherhope/rcloud/app/rclone"
)

// Listen contains the algo for sync
func Listen(d *internal.Directory) {
	var queue chan string = make(chan string)
	var lock bool = false

	go func() {
	runtime:
		for action := range queue {
			if d.RTS {
				lock = true
				switch action {
				case "check":
					message := rclone.Check(d)
					go func() { queue <- message }()
				case "sync":
					message := rclone.Sync(d)
					go func() { queue <- message }()
				case "exit":
					d.SetStatus("exit")
					d.SetChannel(nil)
					os.RemoveAll(d.MakeCachePath(d.Source))
					break runtime
				case "idle":
					d.SetStatus("idle")
				}
				lock = false
			}
		}
	}()

	if d.IsLocal(d.Source) && !d.IsLocal(d.Destination) {
		go runLocalChange(d, lock, queue)
	} else {
		go runRemoteChange(lock, queue)
	}

	d.SetChannel(queue)
}

func runRemoteChange(lock bool, queue chan string) {
	for {
		if !lock {
			//lock = true
			queue <- "check"
		}
		time.Sleep(5 * time.Second)
	}
}

func runLocalChange(d *internal.Directory, lock bool, queue chan string) {
	queue <- "sync"
	for action := range d.CreateMirror(d.Source) {
		if !lock && d.SourceHasChange(action) {
			queue <- "sync"
		}
	}
}
