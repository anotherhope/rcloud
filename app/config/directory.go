package config

import (
	"time"
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
