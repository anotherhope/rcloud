package rclone

import (
	"path"

	"github.com/anotherhope/rcloud/app/internal/repositories"
	"github.com/fsnotify/fsnotify"
)

func CopyOrRemove(r *repositories.Repository, event fsnotify.Event) func() {
	return func() {
		relative := event.Name[len(r.Source):]
		cmd := []string{}
		cmd = append(cmd, r.Args...)
		if event.Op&fsnotify.Remove == fsnotify.Remove || event.Op&fsnotify.Rename == fsnotify.Rename {
			cmd = append(cmd, "delete")
			cmd = append(cmd, path.Join(r.Destination, relative))
		} else {
			cmd = append(cmd, "copyto")
			cmd = append(cmd, path.Join(r.Source, relative))
			cmd = append(cmd, path.Join(r.Destination, relative))
		}

		CreateProcess(r, cmd...)
	}
}

func Sync(r *repositories.Repository) func() {
	return func() {
		SyncFromRepository(r)
	}
}
