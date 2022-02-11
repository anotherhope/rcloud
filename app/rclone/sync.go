package rclone

import (
	"bufio"
	"io"

	"github.com/anotherhope/rcloud/app/config"
)

// Sync execute Rclone sync command run all change
func Sync(d *config.Directory) {
	process := CreateProcess(d.Name, append(
		[]string{
			"sync",
			d.Source,
			d.Destination,
		}, d.Args...)...,
	)

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
}
