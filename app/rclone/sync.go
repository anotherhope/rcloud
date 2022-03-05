package rclone

import (
	"bufio"
	"io"
	"path"

	"github.com/anotherhope/rcloud/app/internal/repositories"
)

func deleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func Sync(rid string) {
	r := repositories.GetRepository(rid)

	cmd := []string{}
	cmd = append(cmd, r.Args...)
	cmd = append(cmd, "sync")
	cmd = append(cmd, ignore(r))
	cmd = append(cmd, path.Join(r.Source))
	cmd = append(cmd, path.Join(r.Destination))

	cmd = deleteEmpty(cmd)

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
			break
		}
	}
}
