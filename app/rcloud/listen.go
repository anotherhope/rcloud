package rcloud

import (
	"github.com/anotherhope/rcloud/app/internal"
)

// Listen contains the algo for sync
func Listen(d *internal.Repository) {
	/*
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
	*/

	if d.IsLocal(d.Source) {
		localDirectory(d)
		//go detectLocalChange(d, lock, queue)
	} else if d.IsRemote(d.Source) {
		remoteDirectory(d)
		//go detectRemoteChange(lock, queue)
	}

	//d.SetChannel(queue)
}

func localDirectory(d *internal.Repository) {

}

func remoteDirectory(d *internal.Repository) {

}

/*
func detectRemoteChange(lock bool, queue chan string) {
	for {
		if !lock {
			queue <- "check"
		}
		time.Sleep(5 * time.Second)
	}
}

func detectLocalChange(d *internal.Repository, lock bool, queue chan string) {
	queue <- "sync"
	for action := range d.CreateMirror(d.Source) {
		if d.RTS && !lock && d.SourceHasChange(action) {
			queue <- "sync"
		}
	}
}
*/
