package rclone

import (
	"bufio"
	"io"
	"path"

	"github.com/anotherhope/rcloud/app/internal/repositories"
	"github.com/fsnotify/fsnotify"
)

func Make(rid string, event fsnotify.Event) func() {
	return func() {
		r := repositories.GetRepository(rid)
		r.SetStatus("sync")
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

		process := CreateProcess(r.Name, cmd...)

		stderr, _ := process.Command.StderrPipe()
		stdout, _ := process.Command.StdoutPipe()
		combined := io.MultiReader(stderr, stdout)
		buf := bufio.NewReader(combined)

		process.Command.Start()
		for {
			_, _, err := buf.ReadLine()
			if err == io.EOF {
				process.Command.Process.Kill()
				process.Command.Process.Wait()
				r.SetStatus("idle")
				break
			}
		}
	}
}
