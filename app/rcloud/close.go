package rcloud

import "github.com/anotherhope/rcloud/app/internal"

func Close(d *internal.Directory) {
	if rc := d.GetChannel(); rc != nil {
		rc <- "exit"
	}
}