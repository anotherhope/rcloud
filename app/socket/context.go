package socket

import (
	"os"
	"path"

	"github.com/anotherhope/rcloud/app/env"
)

func init() {
	os.MkdirAll(path.Dir(env.SocketPath), 0700)
}
