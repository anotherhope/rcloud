package timer

import (
	"fmt"
	"time"
)

type Timer struct {
	ticker *time.Ticker
}

func (t *Timer) Tick() {
	for t := range t.ticker.C {
		fmt.Println("Invoked at ", t)
	}
}

func (t *Timer) Destroy() {
	t.ticker.Stop()
}

func Register(d time.Duration) *Timer {
	return &Timer{
		ticker: time.NewTicker(d),
	}
}
