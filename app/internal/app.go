package internal

import (
	"errors"
	"fmt"
	"os"
	"os/user"
	"path"

	"github.com/spf13/viper"
)

// SocketPath is the path of socket file
var SocketPath string

// CachePath is the path of cache folder
var CachePath string

// User is the current running user
var User *user.User

// Config structure for Rcloud
type Rcloud struct {
	Args         []string
	Repositories []*Directory
}

// App Instance of Rcloud
var App *Rcloud = &Rcloud{
	Args:         []string{},
	Repositories: []*Directory{},
}

func (r *Rcloud) Set(key string, value interface{}) {
	viper.Set(key, value)
	r.Save()
}

// Save write configuration on file
func (r *Rcloud) Save() error {
	return viper.WriteConfig()
}

// Reload the application if config change durint running
/*
func (r *Rcloud) Reload() {
	for _, repository := range App.Repositories {
		fmt.Println(repository.Destination)
		if rc := repository.GetChannel(); rc != nil {
			rc <- "exit"
		}
	}
}
*/

func Load() {
	viper.Unmarshal(App)
}

func init() {
	User, _ = user.Current()
	SocketPath = User.HomeDir + "/.config/rcloud/daemon.sock"
	CachePath = User.HomeDir + "/.config/rcloud/cache"

	if _, err := os.Stat(SocketPath); errors.Is(err, os.ErrNotExist) {
		os.MkdirAll(path.Dir(SocketPath), 0700)
	}

	if _, err := os.Stat(CachePath); errors.Is(err, os.ErrNotExist) {
		os.MkdirAll(CachePath, 0700)
	}

	viper.SetConfigName("rcloud")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.config/rcloud")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			panic(fmt.Errorf("fatal error config file: %w", err))
		}
	}

	viper.SetDefault("config", App.Args)
	viper.SetDefault("repositories", App.Repositories)
	viper.Unmarshal(App)
	viper.SafeWriteConfig()
	viper.WatchConfig()
}
