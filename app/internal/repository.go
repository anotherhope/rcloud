package internal

import (
	"fmt"
	"strings"

	"github.com/anotherhope/rcloud/app/internal/watcher"
)

// Repository is the structure of syncronized folder
type Repository struct {
	Name        string   `mapstructure:"name"`
	Source      string   `mapstructure:"source"`
	Destination string   `mapstructure:"destination"`
	RTS         bool     `mapstructure:"rts"`
	Args        []string `mapstructure:"args"`
	status      string
	watcher     *watcher.Watcher
}

func (d *Repository) Listen() {
	if d.IsLocal(d.Source) {
		d.SetStatus("idle")
		d.watcher, _ = watcher.Register(d.Name, d.Source)
		d.watcher.Handle()

		/*
			go func() {
				var q *queue.Queue
				var pool []fsnotify.Event

				for {
					select {
					case event := <-d.watcher.Change:
						if q == nil {
							q = queue.NewQueue(d.Name)
						} else {
							pool = append(pool, event)
						}
					case <-time.After(1 * time.Second):
						fmt.Println("RUN:", pool)
					}
				}

			}()
		*/
		/*
			d.watcher, d.cache, d.queue = watcher.Register(d.Name, d.Source)
			// New queue initialization.
			productsQueue := queue.NewQueue("NewProducts")
			var jobs []queue.Action
			// Adds jobs to the queue.
			productsQueue.Addactions(jobs)
			// Defines a queue worker, which will execute our queue.
			worker := queue.NewWorker(productsQueue)
			// Execute jobs in queue.
			worker.Execute()
		*/

	} else {
		fmt.Println("Todo")
	}
}

func (d *Repository) Destroy() {
	if d.watcher != nil {
		d.watcher.Destroy()
	}
}

// IsLocal return if path is a local path
func (d *Repository) IsLocal(path string) bool {
	return !d.IsRemote(path)
}

// IsRemote return if path is remote cloud provider
func (d *Repository) IsRemote(path string) bool {
	return strings.Contains(path, ":")
}

// Start synchronization for the current repository
func (d *Repository) Start() error {
	d.RTS = true
	return App.Save()
}

// Stop synchronization for the current repository
func (d *Repository) Stop() error {
	d.RTS = false
	return App.Save()
}

// GetStatus is Getter for Status
func (d *Repository) GetStatus() string {
	return d.status
}

// SetStatus is Setter for Status
func (d *Repository) SetStatus(s string) {
	d.status = s
}
