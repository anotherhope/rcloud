package config

import (
	"crypto/sha256"
	"fmt"
	"io"
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

func createCache(info os.FileInfo, cachePath string, original *os.File) {
	if info.IsDir() {
		os.MkdirAll(cachePath, 0700)
	} else {
		cache, _ := os.Create(cachePath)
		defer cache.Close()

		hash := sha256.New()
		io.Copy(hash, original)

		cache.WriteString(fmt.Sprintf("%x", hash.Sum(nil)))
		os.Chtimes(
			cachePath,
			info.ModTime().Local(),
			info.ModTime().Local(),
		)
	}
}

func walker(d *Directory, pathOfContent string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	relative := pathOfContent[len(d.Source):]
	cachePath := env.CachePath + "/" + d.Name + relative
	cacheStats, _ := os.Stat(cachePath)

	original, _ := os.Open(pathOfContent)
	defer original.Close()

	if cacheStats == nil {
		createCache(info, cachePath, original)
		return nil
	}

	hash := sha256.New()
	io.Copy(hash, original)
	checksum := fmt.Sprintf("%x", hash.Sum(nil))

	if dat, err := os.ReadFile(cachePath); (info.ModTime().Unix() > cacheStats.ModTime().Unix()) ||
		(err == nil && string(dat) != checksum) {
		cache, _ := os.Open(cachePath)
		cache.WriteString(checksum)
		os.Chtimes(
			cachePath,
			info.ModTime().Local(),
			info.ModTime().Local(),
		)
	}

	return nil
}

// SourceHasChange
func (d *Directory) SourceHasChange(pathOfContent string) bool {
	fmt.Println(pathOfContent, d.Destination)

	relative := pathOfContent[len(d.Source):]
	cachePath := env.CachePath + "/" + d.Name + relative
	cacheStats, _ := os.Stat(cachePath)
	originalStats, _ := os.Stat(pathOfContent)

	original, _ := os.Open(pathOfContent)
	defer original.Close()

	hash := sha256.New()
	io.Copy(hash, original)
	checksum := fmt.Sprintf("%x", hash.Sum(nil))

	if dat, err := os.ReadFile(cachePath); (originalStats.ModTime().Unix() > cacheStats.ModTime().Unix()) ||
		(err == nil && string(dat) != checksum) {
		cache, _ := os.Open(cachePath)
		cache.WriteString(checksum)
		os.Chtimes(
			cachePath,
			originalStats.ModTime().Local(),
			originalStats.ModTime().Local(),
		)

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

// HasChange make a mirror of directory to optimize change detect and reduce bandwith comsumption
func (d *Directory) HasChange(pathOfContent string) bool {
	d.SetStatus("check:local")
	err := filepath.Walk(pathOfContent, func(pathOfContent string, info os.FileInfo, err error) error {
		return walker(d, pathOfContent, info, err)
	})
	d.SetStatus("idle")
	return err != nil
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
