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
	"github.com/anotherhope/rcloud/app/internal/repositories"
	"github.com/anotherhope/rcloud/app/rclone"
	"github.com/fsnotify/fsnotify"
)

type Watcher struct {
	rid    string
	notify *fsnotify.Watcher
	cache  *cache.Cache
	change chan fsnotify.Event
}

func (w *Watcher) Queue() {
	var c bool
	var n = make(chan bool)
	var p = make(map[string]func())
	var q *queue.Queue

	for event := range w.change {
		if !c {
			c = true
			p = make(map[string]func())
			q = queue.NewQueue()
		} else {
			if _, ok := p[event.Name]; !ok {
				n <- true
			}
		}

		p[event.Name] = rclone.Make(w.rid, event)

		go func(p map[string]func()) {
			select {
			case <-time.After(250 * time.Millisecond):
				q.Addactions(p)
				queue.NewWorker(q).Execute()
				c = false
			case <-n:
			}
		}(p)
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
	fmt.Println(ignore, pathOfDirectory, repositories.GitIgnore)
	if _, err := os.Stat(path.Join(pathOfDirectory, repositories.GitIgnore)); err == nil {
		ignore.AddFromFS(os.DirFS(pathOfDirectory), repositories.GitIgnore)
	}
	return ignore
}

func Register(rid string, pathOfDirectory string) (*Watcher, error) {
	notify, err := fsnotify.NewWatcher()

	if err != nil {
		return nil, err
	}

	w := &Watcher{
		rid:    rid,
		notify: notify,
		cache:  cache.NewCache(rid, pathOfDirectory),
		change: make(chan fsnotify.Event),
	}

	e := exclude(pathOfDirectory)
	go filepath.Walk(pathOfDirectory, func(currentPath string, info os.FileInfo, err error) error {
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
			if path.Base(event.Name) == repositories.GitIgnore {
				e = exclude(pathOfDirectory)
			}

			if !e.Match(event.Name, true) {
				if event.Op&fsnotify.Remove == fsnotify.Remove || event.Op&fsnotify.Rename == fsnotify.Rename {
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

	go rclone.Sync(rid)

	return w, nil
}
