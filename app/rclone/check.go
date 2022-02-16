package rclone

import (
	"bufio"
	"io"
	"strings"

	"github.com/anotherhope/rcloud/app/internal"
)

// Check execute Rclone check process.Command to detect one change
func Check(d *internal.Directory) string {
	if d.IsLocal(d.Source) && !d.IsLocal(d.Destination) {
		return "idle"
	}

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
	d.SetStatus("check:remote")
	buf := bufio.NewReader(combined)
	count := 0
	for {
		line, _, err := buf.ReadLine()
		if strings.Contains(string(line), "ERROR") {
			count++
		}

		if err == io.EOF {
			process.Command.Process.Kill()
			d.SetStatus("idle")
			break
		}

		if count > 1 {
			process.Command.Process.Kill()
			d.SetStatus("idle")
			return "sync"
		}
	}

	return "idle"
}
