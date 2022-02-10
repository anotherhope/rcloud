package rclone

import (
	"bufio"
	"io"
	"os/exec"
	"strings"

	"github.com/anotherhope/rcloud/app/config"
)

// Check execute Rclone check command to detect one change
func Check(d *config.Directory) bool {
	command := exec.Command("rclone", append(
		[]string{
			"check",
			d.Source,
			d.Destination,
		}, d.Args...)...,
	)
	stderr, _ := command.StderrPipe()
	stdout, _ := command.StdoutPipe()
	combined := io.MultiReader(stderr, stdout)
	command.Start()
	buf := bufio.NewReader(combined)
	count := 0
	for {
		line, _, err := buf.ReadLine()
		if strings.Contains(string(line), "ERROR") {
			count++
		}

		if err == io.EOF {
			command.Process.Kill()
			break
		}

		if count > 1 {
			command.Process.Kill()
			return true
		}
	}

	return false
}
