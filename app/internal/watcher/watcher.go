package watcher

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/aligator/nogo"
	"github.com/anotherhope/rcloud/app/internal/cache"
	"github.com/anotherhope/rcloud/app/internal/queue"
	"github.com/fsnotify/fsnotify"
)

const gitIgnore string = ".gitignore"

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
			case <-time.After(100 * time.Millisecond):
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
			w.notify.Add(currentPath)
			w.cache.Add(currentPath)
		}

		return nil
	})

	go func() {
		for event := range w.notify.Events {
			if path.Base(event.Name) == gitIgnore {
				e = exclude(pathOfDirectory)
			}

			if !e.Match(event.Name, true) {
				if event.Op&fsnotify.Remove == fsnotify.Remove {
					w.notify.Remove(event.Name)
					w.cache.Remove(event.Name)
					w.change <- event
				} else if event.Op&fsnotify.Create == fsnotify.Create {
					sourceInfo, _ := os.Stat(event.Name)
					if sourceInfo.IsDir() {
						sourceParentDirectory := path.Dir(event.Name)

						for strings.Contains(sourceParentDirectory, pathOfDirectory) {
							cacheParentDirectory := w.cache.MakeCachePath(sourceParentDirectory)
							if _, err := os.Stat(sourceParentDirectory); os.IsNotExist(err) {
								os.MkdirAll(cacheParentDirectory, 0700)
							}

							w.notify.Add(sourceParentDirectory)
							sourceParentDirectory = path.Dir(sourceParentDirectory)
						}
					}

					w.notify.Add(event.Name)
					w.cache.Add(event.Name)
					w.change <- event
				} else if event.Op&fsnotify.Write == fsnotify.Write {
					if w.cache.DetectChange(event.Name) {
						w.cache.Update(event.Name)
						w.change <- event
					}
				}
			}

		}
	}()

	go w.Queue()

	return w, nil
}
