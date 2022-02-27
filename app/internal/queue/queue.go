package queue

import (
	"context"
	"log"
	"sync"
)

// Queue holds name, list of actions and context with cancel.
type Queue struct {
	name    string
	actions chan Action
	ctx     context.Context
	cancel  context.CancelFunc
}

// Addactions adds actions to the queue and cancels channel.
func (q *Queue) Addactions(actions []Action) {
	var wg sync.WaitGroup
	wg.Add(len(actions))

	for _, action := range actions {
		// Goroutine which adds Action to the queue.
		go func(action Action) {
			q.AddAction(action)
			wg.Done()
		}(action)
	}

	go func() {
		wg.Wait()
		q.cancel()
	}()
}

// AddAction sends Action to the channel.
func (q *Queue) AddAction(Action Action) {
	q.actions <- Action
	log.Printf("New Action %s added to %s queue", Action.Name, q.name)
}

// NewQueue instantiates new queue.
func NewQueue(name string) *Queue {
	ctx, cancel := context.WithCancel(context.Background())

	return &Queue{
		actions: make(chan Action),
		name:    name,
		ctx:     ctx,
		cancel:  cancel,
	}
}
