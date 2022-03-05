package rclone

import (
	"bufio"
	"fmt"
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

	process := &Process{
		Command: exec.Command("rclone", args...),
	}

	mu.Lock()
	multiton[repository.Name] = &Process{
		Command: exec.Command("rclone", args...),
	}
	mu.Unlock()

	stderr, err := process.Command.StderrPipe()
	fmt.Println(err)
	stdout, err := process.Command.StdoutPipe()
	fmt.Println(err)

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
				repository.SetStatus("")
			}
			break
		} else if err != nil {
			break
		}
	}
}
