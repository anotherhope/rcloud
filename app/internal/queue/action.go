package queue

// Action - holds logic to perform some operations during queue execution.
type Action struct {
	Action func() error // A function that should be executed when the action is running.
}

func (a Action) Run() error {
	err := a.Action()
	if err != nil {
		return err
	}

	return nil
}
