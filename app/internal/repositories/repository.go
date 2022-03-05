package repositories

import (
	"strings"
)

// Repository is the structure of syncronized folder
type Repository struct {
	Name        string   `mapstructure:"name"`
	Source      string   `mapstructure:"source"`
	Destination string   `mapstructure:"destination"`
	RTS         bool     `mapstructure:"rts"`
	Args        []string `mapstructure:"args"`
	status      string
}

// IsSourceLocal return if path is a local path
func (d *Repository) IsSourceLocal() bool {
	return !IsRemote(d.Source)
}

// IsSourceRemote return if path is a local path
func (d *Repository) IsSourceRemote() bool {
	return IsRemote(d.Source)
}

// IsDestinationLocal return if path is a local path
func (d *Repository) IsDestinationLocal() bool {
	return !IsRemote(d.Destination)
}

// IsDestinationRemote return if path is a local path
func (d *Repository) IsDestinationRemote() bool {
	return IsRemote(d.Destination)
}

// IsLocal return if path is a local path
func IsLocal(path string) bool {
	return !IsRemote(path)
}

// IsRemote return if path is remote cloud provider
func IsRemote(path string) bool {
	return strings.Contains(path, ":")
}

// Start synchronization for the current repository
func (d *Repository) Start() {
	d.RTS = true
}

// Stop synchronization for the current repository
func (d *Repository) Stop() {
	d.RTS = false
}

// GetStatus is Getter for Status
func (d *Repository) GetStatus() string {
	return d.status
}

// SetStatus is Setter for Status
func (d *Repository) SetStatus(s string) {
	d.status = s
}
