package main

import (
	"github.com/anotherhope/rcloud/app/cmd"
	"github.com/spf13/viper"
)

func main() {

	cmd.Execute()

	viper.SafeWriteConfig()
	viper.WriteConfig()
}
