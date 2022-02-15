package cmd

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/anotherhope/rcloud/app/config"
	"github.com/anotherhope/rcloud/app/rclone"
	"github.com/anotherhope/rcloud/app/socket"
	"github.com/spf13/cobra"
)

var daemonCmd = &cobra.Command{
	Args:  cobra.ExactArgs(0),
	Use:   "daemon",
	Short: "Daemon management",
	RunE: func(cmd *cobra.Command, args []string) error {
		var exit = make(chan os.Signal, 1)
		signal.Notify(exit, os.Interrupt)

		go socket.Server()
		for _, repository := range config.Load().Repositories {
			go rclone.Daemon(repository)
		}

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
