package timer

import (
	"time"

	"github.com/anotherhope/rcloud/app/rclone"
)

type Timer struct {
	ticker *time.Ticker
	rid    string
}

func (t *Timer) Tick() {
	rclone.Sync(t.rid)
	var lock bool = false
	for range t.ticker.C {
		if !lock {
			lock = true
			rclone.Sync(t.rid)
			lock = false
		}
	}
}

func (t *Timer) Destroy() {
	t.ticker.Stop()
}

func Register(rid string, d time.Duration) *Timer {
	t := &Timer{
		ticker: time.NewTicker(d),
		rid:    rid,
	}

	go t.Tick()

	return t
}
