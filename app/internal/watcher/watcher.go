package watcher

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/aligator/nogo"
	"github.com/anotherhope/rcloud/app/internal/cache"
	"github.com/anotherhope/rcloud/app/internal/queue"
	"github.com/fsnotify/fsnotify"
)

type Watcher struct {
	notify *fsnotify.Watcher
	cache  *cache.Cache
	change chan fsnotify.Event
}

func (w *Watcher) Queue() {

	var canceled = make(chan bool)
	var gotChange bool
	var pools = make([]queue.Action, 0)
	var q *queue.Queue

	for event := range w.change {
		if !gotChange {
			gotChange = true
			q = queue.NewQueue()
		} else {
			pools = append(pools, queue.Action{
				Action: func() error {
					return nil
				},
			})
			canceled <- true
		}

		go func(event fsnotify.Event) {
			select {
			case <-time.After(5 * time.Second):
				fmt.Println("do it", event)
				q.Addactions(pools)
				w := queue.NewWorker(q)
				w.Execute()
				pools = make([]queue.Action, 0)
				gotChange = false
			case <-canceled:
				fmt.Println("abort", event)
			}
		}(event)
	}
}

func (w *Watcher) Destroy() {
	if w.cache != nil {
		w.cache.Remove()
		w.cache = nil
	}

	if w.notify != nil {
		w.notify.Close()
		w.notify = nil
	}
}

const gitIgnore string = ".gitignore"

func exclude(pathOfDirectory string) *nogo.NoGo {
	ignore := nogo.New(nogo.DotGitRule)
	ignore.AddFromFS(os.DirFS(pathOfDirectory), gitIgnore)
	return ignore
}

func Register(idOfRepository string, pathOfDirectory string) (*Watcher, error) {
	notify, err := fsnotify.NewWatcher()

	if err != nil {
		return nil, err
	}

	w := &Watcher{
		notify: notify,
		cache:  cache.NewCache(idOfRepository, pathOfDirectory),
		change: make(chan fsnotify.Event),
	}

	e := exclude(pathOfDirectory)

	filepath.Walk(pathOfDirectory, func(currentPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !e.Match(currentPath, true) {
			w.cache.DetectChange(currentPath)
			w.notify.Add(currentPath)
		}

		return nil
	})

	go func() {
		for event := range w.notify.Events {
			if event.Op&fsnotify.Chmod == fsnotify.Chmod {
				continue
			} else if event.Op&fsnotify.Remove == fsnotify.Remove {
				if path.Base(event.Name) == gitIgnore {
					w.cache.Remove(event.Name)
					e = nil
				}

				if !e.Match(event.Name, true) {
					w.notify.Remove(event.Name)
					w.change <- event
				}
			} else {
				if path.Base(event.Name) == gitIgnore {
					e = exclude(pathOfDirectory)
				}

				if !e.Match(event.Name, true) {
					if event.Op&fsnotify.Create == fsnotify.Create {
						w.notify.Add(event.Name)
					}

					if w.cache.DetectChange(event.Name) {
						w.change <- event
					}
				}
			}
		}
	}()

	go w.Queue()

	return w, nil
}

//
