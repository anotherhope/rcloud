package rclone

import (
	"os/exec"
	"sync"
)

// Process is the structure of a subprocess
type Process struct {
	Command *exec.Cmd
	Type    string
}

var mu = sync.Mutex{}
var multiton = map[string]*Process{}

// CreateProcess can create a new process for rclone
func CreateProcess(repositoryName string, args ...string) *Process {
	mu.Lock()
	defer mu.Unlock()

	multiton[repositoryName] = &Process{
		Command: exec.Command("rclone", args...),
		Type:    args[0],
	}

	return multiton[repositoryName]
}
