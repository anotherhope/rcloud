package rclone

import (
	"os"
	"os/exec"

	"github.com/anotherhope/rcloud/app/internal"
)

const ignore = "/.gitignore"

func gitIgnore(d *internal.Repository) []string {
	ignores := make([]string, 0)

	if d.IsLocal(d.Source) {
		if _, err := os.Stat(d.Source + ignore); err != nil {
			ignores = append(ignores, "--exclude=\""+d.Source+ignore+"\"")
		}
	}

	if d.IsRemote(d.Source) {
		cmd := exec.Command("rclone", "copyto", d.Source+ignore, d.Destination+ignore)
		if err := cmd.Wait(); err == nil {
			ignores = append(ignores, "--exclude=\""+d.Destination+ignore+"\"")
		}
	}

	return ignores
}
