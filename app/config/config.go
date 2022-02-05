package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Repositories []*Directory
}

type Directory struct {
	Name        string        `mapstructure:"name"`
	Source      string        `mapstructure:"source"`
	Destination string        `mapstructure:"destination"`
	Watch       time.Duration `mapstructure:"watch"`
}

func (d *Directory) Status() string {

	return ""
}

var instance *Config = &Config{
	Repositories: []*Directory{},
}

func Get() *Config {
	return instance
}

func Set(key string, value interface{}) {
	viper.Set(key, value)
	viper.WriteConfig()
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

	viper.Unmarshal(instance)

	viper.SafeWriteConfig()
	viper.WriteConfig()
}
