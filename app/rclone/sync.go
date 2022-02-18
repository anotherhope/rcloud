package rclone

import (
	"bufio"
	"fmt"
	"io"

	"github.com/anotherhope/rcloud/app/internal"
)

// Sync execute Rclone sync command run all change
func Sync(d *internal.Directory) string {
	d.SetStatus("sync")
	cmd := []string{"sync", d.Source, d.Destination}
	cmd = append(cmd, d.Args...)
	cmd = append(cmd, gitIgnore(d)...)

	fmt.Println(d.Name, cmd) //
	process := CreateProcess(d.Name, cmd...)

	stderr, _ := process.Command.StderrPipe()
	stdout, _ := process.Command.StdoutPipe()
	combined := io.MultiReader(stderr, stdout)
	process.Command.Start()
	buf := bufio.NewReader(combined)
	for {
		_, _, err := buf.ReadLine()
		if err == io.EOF {
			process.Command.Process.Kill()
			break
		}
	}

	return "idle"
}
