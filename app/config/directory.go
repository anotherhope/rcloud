package config

import (
	"crypto/sha256"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/anotherhope/rcloud/app/env"
	"github.com/fsnotify/fsnotify"
)

// Directory is the structure of syncronized folder
type Directory struct {
	Name        string        `mapstructure:"name"`
	Source      string        `mapstructure:"source"`
	Destination string        `mapstructure:"destination"`
	Watch       time.Duration `mapstructure:"watch"`
	RTS         bool          `mapstructure:"rts"`
	Args        []string      `mapstructure:"args"`
	status      string
}

func (d *Directory) makeCachePath(pathOfContent string) string {
	relative := pathOfContent[len(d.Source):]
	return env.CachePath + "/" + d.Name + relative
}

func updateTimes(pathOfCache string, info fs.FileInfo) {
	os.Chtimes(
		pathOfCache,
		info.ModTime().Local(),
		info.ModTime().Local(),
	)
}

func makeHash(original io.Reader) string {
	hash := sha256.New()
	io.Copy(hash, original)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func walker(d *Directory, pathOfContent string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	cachePath := d.makeCachePath(pathOfContent)

	if info.IsDir() {
		os.MkdirAll(cachePath, 0700)
		updateTimes(cachePath, info)
		return nil
	}

	original, _ := os.Open(pathOfContent)
	defer original.Close()

	cache, _ := os.OpenFile(cachePath, os.O_RDWR|os.O_CREATE, 0600)
	defer cache.Close()
	cache.WriteString(makeHash(original))
	updateTimes(cachePath, info)

	return nil
}

// SourceHasChange
func (d *Directory) SourceHasChange(pathOfContent string) bool {

	cachePath := d.makeCachePath(pathOfContent)
	original, _ := os.Open(pathOfContent)
	defer original.Close()

	checksum := makeHash(original)
	if dat, err := os.ReadFile(cachePath); err == nil && string(dat) != checksum {
		os.Truncate(cachePath, 0)
		cache, _ := os.OpenFile(cachePath, os.O_WRONLY, 0700)
		defer cache.Close()
		cache.WriteString(checksum)
		originalStats, _ := original.Stat()
		updateTimes(cache.Name(), originalStats)
		return true
	}

	return false
}

// CreateMirror make a mirror of directory to optimize change detect and reduce bandwith comsumption
func (d *Directory) CreateMirror(pathOfContent string) chan string {
	watcher, _ := fsnotify.NewWatcher()
	var action chan string = make(chan string)

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}

				if event.Op&fsnotify.Write == fsnotify.Write {
					action <- event.Name
				}

				if event.Op&fsnotify.Create == fsnotify.Create {
					watcher.Add(event.Name)
					action <- event.Name
				}

				if event.Op&fsnotify.Remove == fsnotify.Remove {
					watcher.Remove(event.Name)
					action <- event.Name
				}

				if event.Op&fsnotify.Rename == fsnotify.Rename {
					action <- event.Name
				}

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	filepath.Walk(pathOfContent, func(currentPath string, info os.FileInfo, err error) error {
		watcher.Add(currentPath)
		return walker(d, currentPath, info, err)
	})

	return action
}

// IsLocal return if path is a local path
func (d *Directory) IsLocal(path string) bool {
	return !d.IsRemote(path)
}

// IsRemote return if path is remote cloud provider
func (d *Directory) IsRemote(path string) bool {
	return strings.Contains(path, ":")
}

// Start synchronization for the current directory
func (d *Directory) Start() error {
	d.RTS = true
	return Save()
}

// Stop synchronization for the current directory
func (d *Directory) Stop() error {
	d.RTS = false
	return Save()
}

// GetStatus is Getter for Status
func (d *Directory) GetStatus() string {
	return d.status
}

// SetStatus is Setter for Status
func (d *Directory) SetStatus(s string) {
	d.status = s
}
