package rclone

import (
	"bufio"
	"io"
	"os/exec"

	"github.com/anotherhope/rcloud/app/config"
)

// Sync execute Rclone sync command run all change
func Sync(d *config.Directory) {
	command := exec.Command("rclone", append(
		[]string{
			"sync",
			d.Source,
			d.Destination,
		}, d.Args...)...,
	)
	stderr, _ := command.StderrPipe()
	stdout, _ := command.StdoutPipe()
	combined := io.MultiReader(stderr, stdout)
	command.Start()
	buf := bufio.NewReader(combined)
	for {
		_, _, err := buf.ReadLine()
		if err == io.EOF {
			command.Process.Kill()
			break
		}
	}
}
