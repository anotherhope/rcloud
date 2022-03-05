package rclone

import (
	"bufio"
	"io"
	"path"

	"github.com/anotherhope/rcloud/app/internal/repositories"
)

func Sync(rid string) {
	r := repositories.GetRepository(rid)

	cmd := []string{}
	cmd = append(cmd, r.Args...)
	cmd = append(cmd, "sync")
	cmd = append(cmd, ignore(r))
	cmd = append(cmd, path.Join(r.Source))
	cmd = append(cmd, path.Join(r.Destination))

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
			break
		}
	}
}
