package watcher

import (
	"os"
	"path"
	"path/filepath"

	"github.com/aligator/nogo"
	"github.com/anotherhope/rcloud/app/internal/cache"
	"github.com/anotherhope/rcloud/app/internal/queue"
	"github.com/fsnotify/fsnotify"
)

type Watcher struct {
	watcher *fsnotify.Watcher
	cache   *cache.Cache
	queue   *queue.Queue
}

func (w *Watcher) Destroy() {
	if w.cache != nil {
		w.cache.Remove()
		w.cache = nil
	}

	if w.watcher != nil {
		w.watcher.Close()
		w.watcher = nil
	}
}

const gitIgnore string = ".gitignore"

func exclude(pathOfDirectory string) *nogo.NoGo {
	ignore := nogo.New(nogo.DotGitRule)
	ignore.AddFromFS(os.DirFS(pathOfDirectory), gitIgnore)
	return ignore
}

/*
func Remove(pathOfDirectory string) {

}
*/

func Register(idOfRepository string, pathOfDirectory string) (*fsnotify.Watcher, *cache.Cache, chan fsnotify.Event) {
	a := make(chan fsnotify.Event)
	w, _ := fsnotify.NewWatcher()
	e := exclude(pathOfDirectory)
	c := cache.Make(idOfRepository, pathOfDirectory)

	filepath.Walk(pathOfDirectory, func(currentPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !e.Match(currentPath, true) {
			c.DetectChange(currentPath)
			w.Add(currentPath)
		}

		return nil
	})

	go func() {
		for event := range w.Events {
			if event.Op&fsnotify.Remove == fsnotify.Remove {
				w.Remove(event.Name)
				if path.Base(event.Name) == gitIgnore {
					c.Remove(event.Name)
					e = nil
				}
			} else {
				if path.Base(event.Name) == gitIgnore {
					e = exclude(pathOfDirectory)
				}

				if !e.Match(event.Name, true) {
					if event.Op&fsnotify.Create == fsnotify.Create {
						w.Add(event.Name)
					}

					if c.DetectChange(event.Name) {
						a <- event
					}
				}
			}
		}
	}()

	return w, c, a
}
