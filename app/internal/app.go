package internal

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config structure for rcloud
type Rcloud struct {
	Args         []string
	Repositories []*Directory
}

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

func init() {
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
	viper.WriteConfig()
}
