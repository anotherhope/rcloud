package rclone

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/anotherhope/rcloud/app/internal"
)

func gitIgnore(d *internal.Directory) []string {
	ignores := make([]string, 0)

	if d.IsLocal(d.Source) {
		if _, err := os.Stat(d.Source + "/.gitignore"); err != nil {
			ignores = append(ignores, "--exclude=\""+d.Source+"/.gitignore"+"\"")
		}
	}

	if d.IsRemote(d.Source) {
		cmd := exec.Command("rclone", "copyto", d.Source+"/.gitignore", d.Destination+"/.gitignore")
		if err := cmd.Wait(); err == nil {
			fmt.Println("ok")
			ignores = append(ignores, "--exclude=\""+d.Destination+"/.gitignore"+"\"")
		}
	}

	return ignores
}
