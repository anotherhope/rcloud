package internal

import (
	"fmt"

	"github.com/spf13/viper"
)

// App Instance of Rcloud
var App *Rcloud

// Rcloud structure is configuration file for Rcloud
type Rcloud struct {
	Args         []string
	Repositories []*Repository
}

func (r *Rcloud) Load() {
	for _, repository := range App.Repositories {
		repository.Destroy()
	}
	UpdateConfig()
	for _, repository := range App.Repositories {
		repository.Listen()
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
		Args:         []string{},
		Repositories: []*Repository{},
	}

	viper.Unmarshal(App)
}

func init() {
	viper.SetConfigName("rcloud")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.config/rcloud")

	UpdateConfig()

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
