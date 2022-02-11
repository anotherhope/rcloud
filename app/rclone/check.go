package rclone

import (
	"bufio"
	"io"
	"strings"

	"github.com/anotherhope/rcloud/app/config"
)

// Check execute Rclone check process.Command to detect one change
func Check(d *config.Directory) bool {
	process := CreateProcess(d.Name, append(
		[]string{
			"check",
			d.Source,
			d.Destination,
			"--fast-list",
			"--checkers=1",
		}, d.Args...)...,
	)

	stderr, _ := process.Command.StderrPipe()
	stdout, _ := process.Command.StdoutPipe()
	combined := io.MultiReader(stderr, stdout)
	process.Command.Start()
	buf := bufio.NewReader(combined)
	count := 0
	for {
		line, _, err := buf.ReadLine()
		if strings.Contains(string(line), "ERROR") {
			count++
		}

		if err == io.EOF {
			process.Command.Process.Kill()
			break
		}

		if count > 1 {
			process.Command.Process.Kill()
			return true
		}
	}

	return false
}
