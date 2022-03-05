package config

import (
	"fmt"
	"time"

	"github.com/anotherhope/rcloud/app/internal/repositories"
	"github.com/anotherhope/rcloud/app/internal/system"
	"github.com/anotherhope/rcloud/app/internal/timer"
	"github.com/anotherhope/rcloud/app/internal/watcher"
	"github.com/spf13/viper"
)

// App Instance of Rcloud
var App *Rcloud

// Rcloud structure is configuration file for Rcloud
type Rcloud struct {
	Repositories *[]*repositories.Repository
	Watcher      map[string]*watcher.Watcher
	Timer        map[string]*timer.Timer
}

func (r *Rcloud) Load() {
	for _, watcher := range App.Watcher {
		watcher.Destroy()
	}
	for _, timer := range App.Timer {
		timer.Destroy()
	}
	UpdateConfig()
	for _, r := range repositories.Repositories {
		if r.RTS {
			if r.IsSourceLocal() {
				App.Watcher[r.Name], _ = watcher.Register(r.Name, r.Source)
			} else {
				App.Timer[r.Name] = timer.Register(r.Name, 1*time.Minute)
			}
		} else {
			r.SetStatus("aaaaa")
		}
	}
}

func (r *Rcloud) Set(key string, value interface{}) {
	viper.Set(key, value)
	r.Save()
}

// Save write configuration on file
func (r *Rcloud) Save() error {
	return viper.WriteConfig()
}

func UpdateConfig() {
	viper.ReadInConfig()

	App = &Rcloud{
		Repositories: &repositories.Repositories,
		Watcher:      map[string]*watcher.Watcher{},
		Timer:        map[string]*timer.Timer{},
	}

	viper.Unmarshal(App)
}

func init() {

	viper.SetConfigName("rcloud")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(system.User.HomeDir + "/.config/rcloud")

	UpdateConfig()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			panic(fmt.Errorf("fatal error config file: %w", err))
		}
	}

	//viper.SetDefault("config", App.Args)
	//viper.SetDefault("repositories", App.Repositories)

	viper.SafeWriteConfig()
	viper.WatchConfig()
}
