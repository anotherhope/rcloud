package queue

import "log"

// Action - holds logic to perform some operations during queue execution.
type Action struct {
	Name   string
	Action func() error // A function that should be executed when the action is running.
}

func (a Action) Run() error {
	log.Printf("Job running: %s", a.Name)

	err := a.Action()
	if err != nil {
		return err
	}

	return nil
}
