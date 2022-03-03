package queue

// Worker responsible for queue serving.
type Worker struct {
	Queue *Queue
}

// NewWorker initializes a new Worker.
func NewWorker(queue *Queue) *Worker {
	return &Worker{
		Queue: queue,
	}
}

// Execute processes actions from the queue (actions channel).
func (w *Worker) Execute() bool {
	for {
		select {
		case <-w.Queue.ctx.Done():
			return true
		case action := <-w.Queue.actions:
			action.Execute()
		}
	}
}
