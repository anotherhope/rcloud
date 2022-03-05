package config

import (
	"fmt"
	"os"

	"github.com/anotherhope/rcloud/app/internal/repositories"
	"github.com/anotherhope/rcloud/app/internal/watcher"
	"github.com/spf13/viper"
)

// App Instance of Rcloud
var App *Rcloud

// Rcloud structure is configuration file for Rcloud
type Rcloud struct {
	//Args         []string
	Repositories *[]*repositories.Repository
	Watcher      map[string]*watcher.Watcher
}

func (r *Rcloud) Load() {
	for _, watcher := range App.Watcher {
		watcher.Destroy()
	}
	UpdateConfig()
	for _, r := range repositories.Repositories {
		if r.IsSourceLocal() {
			App.Watcher[r.Name], _ = watcher.Register(r.Name, r.Source)
			go App.Watcher[r.Name].Status(r)
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
	}

	viper.Unmarshal(App.Repositories)
}

func init() {
	home, _ := os.UserHomeDir()

	viper.SetConfigName("rcloud")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(home + "/.config/rcloud")

	UpdateConfig()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			panic(fmt.Errorf("fatal error config file: %w", err))
		}
	}

	//viper.SetDefault("config", App.Args)
	//viper.SetDefault("repositories", App.Repositories)

	viper.Unmarshal(App)
	viper.SafeWriteConfig()
	viper.WatchConfig()
}
