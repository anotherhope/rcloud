package rclone

import (
	"bufio"
	"io"
	"os/exec"
	"strings"
)

// Check execute Rclone check command to detect one change
func Check(s string, d string) bool {
	command := exec.Command("rclone", "check", s, d, "--transfers=1")
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

// Sync execute Rclone sync command run all change
func Sync(s string, d string) {
	command := exec.Command("rclone", "sync", s, d)
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
