package config

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/anotherhope/rcloud/app/env"
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

// HasChange make a mirror of folder to optimize change detect and reduce bandwith comsumption
func (d *Directory) HasChange(pathOfContent string) bool {
	err := filepath.Walk(pathOfContent, func(pathOfContent string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relative := pathOfContent[len(d.Source):]
		cachePath := env.CachePath + "/" + d.Name + relative

		if info.IsDir() {
			os.MkdirAll(cachePath, 0700)
			return nil
		}

		cacheStats, _ := os.Stat(cachePath)
		original, _ := os.Open(pathOfContent)
		defer original.Close()

		hash := sha256.New()
		io.Copy(hash, original)
		checksum := fmt.Sprintf("%x", hash.Sum(nil))

		if cacheStats == nil {
			cache, _ := os.Create(cachePath)
			defer cache.Close()

			cache.WriteString(checksum)
			os.Chtimes(
				cachePath,
				info.ModTime().Local(),
				info.ModTime().Local(),
			)
		} else {
			if cacheStats.ModTime().Unix() < info.ModTime().Unix() && false {
				os.Chtimes(
					cachePath,
					info.ModTime().Local(),
					info.ModTime().Local(),
				)
			} else if dat, err := os.ReadFile(cachePath); err == nil && string(dat) != checksum {
				cache, _ := os.Open(cachePath)
				cache.WriteString(checksum)
			}
		}

		return nil
	})

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
