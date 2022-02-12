package rclone

import (
	"os/exec"
	"sync"
)

type Process struct {
	Command *exec.Cmd
	Type    string
}

var mu = sync.Mutex{}
var multiton = map[string]*Process{}

func CreateProcess(directoryName string, args ...string) *Process {
	mu.Lock()
	defer mu.Unlock()

	multiton[directoryName] = &Process{
		Command: exec.Command("rclone", args...),
		Type:    args[0],
	}

	return multiton[directoryName]
}
