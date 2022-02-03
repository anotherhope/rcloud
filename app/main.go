package main

import (
	"fmt"

	"github.com/anotherhope/rcloud/app/cmd"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("rcloud")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.config/rcloud")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			panic(fmt.Errorf("fatal error config file: %w", err))
		}
	}

	conf := LoadConfig()

	fmt.Println("conf", conf)

	cmd.Execute()
}
