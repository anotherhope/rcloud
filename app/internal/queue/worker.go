package queue

import "log"

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

// Execute processes jobs from the queue (jobs channel).
func (w *Worker) Execute() bool {
	for {
		select {
		case <-w.Queue.ctx.Done():
			return true
		case job := <-w.Queue.actions:
			err := job.Run()
			if err != nil {
				log.Print(err)
				continue
			}
		}
	}
}
