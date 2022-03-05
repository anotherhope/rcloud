package timer

import (
	"fmt"
	"time"

	"github.com/anotherhope/rcloud/app/rclone"
)

type Timer struct {
	ticker *time.Ticker
	rid    string
}

func (t *Timer) Tick() {
	var lock bool = false
	for range t.ticker.C {
		if !lock {
			fmt.Println("start")
			lock = true
			rclone.Sync(t.rid)
			lock = false
			fmt.Println("stop")
		}
	}
}

func (t *Timer) Destroy() {
	t.ticker.Stop()
}

func Register(rid string, d time.Duration) *Timer {
	return &Timer{
		ticker: time.NewTicker(d),
		rid:    rid,
	}
}
