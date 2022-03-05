package rclone

import (
	"bufio"
	"io"
	"os/exec"
	"sync"

	"github.com/anotherhope/rcloud/app/internal/repositories"
)

// Process is the structure of a subprocess
type Process struct {
	Command *exec.Cmd
}

var mu = sync.Mutex{}
var multiton = map[string]*Process{}

// CreateProcess can create a new process for rclone
func CreateProcess(repository *repositories.Repository, args ...string) {
	repository.SetStatus("sync")
	process := &Process{
		Command: exec.Command("rclone", args...),
	}

	mu.Lock()
	multiton[repository.Name] = &Process{
		Command: exec.Command("rclone", args...),
	}
	mu.Unlock()

	stderr, _ := process.Command.StderrPipe()
	stdout, _ := process.Command.StdoutPipe()

	combined := io.MultiReader(stderr, stdout)
	buf := bufio.NewReader(combined)

	process.Command.Start()
	for {
		if buf == nil {
			break
		}
		_, _, err := buf.ReadLine()
		if err == io.EOF {
			process.Command.Process.Kill()
			process.Command.Process.Wait()
			mu.Lock()
			multiton[repository.Name] = nil
			mu.Unlock()
			if repository.RTS {
				repository.SetStatus("idle")
			} else {
				repository.SetStatus("bbbbbb")
			}

			stderr.Close()
			stdout.Close()

			break
		} else if err != nil {
			break
		}
	}
}
