package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config structure for rcloud
type Config struct {
	Args         []string
	Repositories []*Directory
}

var instance = &Config{
	Args:         []string{"--fast-list"},
	Repositories: []*Directory{},
}

// Load value in configuration file
func Load() *Config {
	return instance
}

// Set value in confuration file and save behind
func Set(key string, value interface{}) {
	viper.Set(key, value)
	Save()
}

// Save write configuration on file
func Save() error {
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

	viper.SetDefault("config", instance.Args)
	viper.SetDefault("repositories", instance.Repositories)

	viper.Unmarshal(instance)

	viper.SafeWriteConfig()
	viper.WriteConfig()
}
