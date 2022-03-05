package rclone

import (
	"os/exec"
	"sync"
)

// Process is the structure of a subprocess
type Process struct {
	Command *exec.Cmd
}

var mu = sync.Mutex{}
var multiton = map[string]*Process{}

// CreateProcess can create a new process for rclone
func CreateProcess(repositoryName string, args ...string) *Process {
	mu.Lock()
	multiton[repositoryName] = &Process{
		Command: exec.Command("rclone", args...),
	}
	mu.Unlock()
	return multiton[repositoryName]
}

func RemoveProcess(repositoryName string) {
	mu.Lock()
	if process, ok := multiton[repositoryName]; ok {
		process.Command.Process.Kill()
		process.Command.Process.Wait()
		multiton[repositoryName] = nil
	}
	mu.Unlock()
}
