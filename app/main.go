package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/anotherhope/rcloud/app/cmd"
)

func init() {
	_, err := exec.Command("which", "rclone").Output()
	if err != nil {
		fmt.Println("Please install rclone before use rcloud")
		os.Exit(1)
	}
}

func main() {
	cmd.Execute()
}
