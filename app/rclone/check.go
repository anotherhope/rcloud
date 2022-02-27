package rclone

import (
	"bufio"
	"io"
	"strings"

	"github.com/anotherhope/rcloud/app/internal"
)

// Check execute Rclone check cmd to detect one change
func Check(d *internal.Repository) string {
	if d.IsLocal(d.Source) {
		return "idle"
	}

	d.SetStatus("check")

	cmd := []string{"check", d.Source, d.Destination, "--fast-list", "--checkers=1"}
	cmd = append(cmd, d.Args...)
	cmd = append(cmd, gitIgnore(d)...)

	process := CreateProcess(d.Name, cmd...)

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
