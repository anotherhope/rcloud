package cmd

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/anotherhope/rcloud/app/internal"
	"github.com/anotherhope/rcloud/app/rclone"
	"github.com/anotherhope/rcloud/app/rcloud"
	"github.com/anotherhope/rcloud/app/socket"
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

		go socket.Server()
		for _, repository := range internal.App.Repositories {
			rcloud.Listen(repository)
		}

		viper.OnConfigChange(func(e fsnotify.Event) {
			for _, repository := range internal.App.Repositories {
				rcloud.Close(repository)
			}
			internal.Load()
			for _, repository := range internal.App.Repositories {
				rcloud.Listen(repository)
			}
		})

		<-exit
		rclone.Kill()
		fmt.Println()
		return nil
	},
	DisableFlagsInUseLine: true,
}

func init() {
	rootCmd.AddCommand(daemonCmd)
}
