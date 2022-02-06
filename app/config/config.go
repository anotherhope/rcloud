package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// Config is the struct of the config file for unmarshal
type Config struct {
	Repositories []*Directory
}

// Directory is the structure of syncronized folder
type Directory struct {
	Name        string        `mapstructure:"name"`
	Source      string        `mapstructure:"source"`
	Destination string        `mapstructure:"destination"`
	Watch       time.Duration `mapstructure:"watch"`
}

var instance *Config = &Config{
	Repositories: []*Directory{},
}

// Status can display statement of a Directory
func (d *Directory) Status() string {

	return ""
}

// Get configuration instance
func Get() *Config {
	return instance
}

// Set value in confuration file and save behind
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
