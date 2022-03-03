package queue

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
)

// Action - holds logic to perform some operations during queue execution.
type Action struct {
	Rid   string
	Event fsnotify.Event
}

func (a Action) Execute() {
	fmt.Println(a)
}
