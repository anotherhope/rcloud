package rclone

import (
	"os"
	"os/exec"
	"path"

	"github.com/anotherhope/rcloud/app/internal/repositories"
	"github.com/anotherhope/rcloud/app/internal/system"
)

func ignore(d *repositories.Repository) string {
	if repositories.IsLocal(d.Source) {
		src := path.Join(d.Source, repositories.GitIgnore)
		if _, err := os.Stat(src); err == nil {
			return "--exclude=\"" + src + "\""
		}
	} else if repositories.IsRemote(d.Source) && repositories.IsLocal(d.Destination) {
		src := path.Join(d.Source, repositories.GitIgnore)
		dst := path.Join(d.Destination, repositories.GitIgnore)

		cmd := exec.Command("rclone", "copyto", src, dst)
		if err := cmd.Wait(); err == nil {
			return "--exclude=\"" + dst + "\""
		}
	} else if repositories.IsRemote(d.Source) && repositories.IsRemote(d.Destination) {
		src := path.Join(d.Source, repositories.GitIgnore)
		lcl := path.Join(system.User.HomeDir, ".config/rcloud/tmp", d.Name, repositories.GitIgnore)
		cmd := exec.Command("rclone", "copyto", src, lcl)
		if err := cmd.Wait(); err == nil {
			return "--exclude=\"" + lcl + "\""
		}
	}

	return ""
}
