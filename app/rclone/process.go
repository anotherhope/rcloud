package rclone

import (
	"os/exec"
)

type Process struct {
	Command *exec.Cmd
	Type    string
}

var multiton = map[string]*Process{}

func CreateProcess(directoryName string, args ...string) *Process {
	multiton[directoryName] = &Process{
		Command: exec.Command("rclone", args...),
		Type:    args[0],
	}
	return multiton[directoryName]
}
