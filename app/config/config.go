package config

import (
	"fmt"

	"github.com/spf13/viper"
)

/*
// Config is the struct of the config file for unmarshal
type Config struct {
	Repositories []*Directory
}

/*
var instance *Config = &Config{
	Repositories: []*Directory{},
}

/*
// Get configuration instance
func Get() *Config {
	return instance
}
*/

//  Cast value from configuration file
func Cast(k string, i interface{}) {
	viper.UnmarshalKey(k, &i)
}

//  Get value in configuration file
func Get(key string) interface{} {
	return viper.Get(key)
}

// Set value in confuration file and save behind
func Set(key string, value interface{}) {
	viper.Set(key, value)
	Save()
}

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

	//viper.Unmarshal(instance)

	viper.SafeWriteConfig()
	viper.WriteConfig()
}
