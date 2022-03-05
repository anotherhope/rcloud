package cmd

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/anotherhope/rcloud/app/internal/cache"
	"github.com/anotherhope/rcloud/app/internal/config"
	"github.com/anotherhope/rcloud/app/internal/socket"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var daemonCmd = &cobra.Command{
	Args:  cobra.ExactArgs(0),
	Use:   "daemon",
	Short: "Start Daemon service",
	RunE: func(cmd *cobra.Command, args []string) error {
		var exit = make(chan os.Signal, 1)
		signal.Notify(exit, os.Interrupt)

		os.RemoveAll(cache.CachePath)

		go socket.Server()
		go config.App.Load()

		viper.OnConfigChange(func(e fsnotify.Event) {
			go config.App.Load()
		})

		<-exit
		socket.Stop()
		fmt.Println()
		return nil
	},
	DisableFlagsInUseLine: true,
}

func init() {
	rootCmd.AddCommand(daemonCmd)
}
