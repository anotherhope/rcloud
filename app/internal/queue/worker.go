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
		// if context was canceled.
		case <-w.Queue.ctx.Done():
			log.Printf("Work done in queue: %s!", w.Queue.ctx.Err())
			return true
		// if job received.
		case job := <-w.Queue.actions:
			err := job.Run()
			if err != nil {
				log.Print(err)
				continue
			}
		}
	}
}
