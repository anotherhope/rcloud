package rclone

import (
	"path"

	"github.com/anotherhope/rcloud/app/internal/repositories"
)

func deleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func Sync(rid string) {
	//fmt.Println(rid, "sync")
	r := repositories.GetRepository(rid)
	cmd := []string{}
	cmd = append(cmd, r.Args...)
	cmd = append(cmd, "sync")
	cmd = append(cmd, ignore(r))
	cmd = append(cmd, path.Join(r.Source))
	cmd = append(cmd, path.Join(r.Destination))
	cmd = deleteEmpty(cmd)
	CreateProcess(r, cmd...)
}
