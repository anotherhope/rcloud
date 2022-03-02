package queue

import "github.com/fsnotify/fsnotify"

// Action - holds logic to perform some operations during queue execution.
type Action struct {
	Event  fsnotify.Event
	Action func() error // A function that should be executed when the action is running.
}

func (a Action) Run() error {
	err := a.Action()
	if err != nil {
		return err
	}

	return nil
}
