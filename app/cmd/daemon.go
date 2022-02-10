package cmd

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/anotherhope/rcloud/app/config"
	"github.com/anotherhope/rcloud/app/rclone"
	"github.com/spf13/cobra"
)

var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Daemon management",
}

var daemonInstall = &cobra.Command{
	Args:  cobra.ExactArgs(0),
	Use:   "install",
	Short: "Install daemon",
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO
		fmt.Println("install")
		return nil
	},
}

var daemonRemove = &cobra.Command{
	Args:  cobra.ExactArgs(0),
	Use:   "remove",
	Short: "Remove daemon",
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO
		fmt.Println("remove")
		return nil
	},
}

var daemonStart = &cobra.Command{
	Args:  cobra.ExactArgs(0),
	Use:   "start",
	Short: "Run daemon in start mode",
	RunE: func(cmd *cobra.Command, args []string) error {
		var exit = make(chan os.Signal, 1)
		signal.Notify(exit, os.Interrupt)

		for _, repository := range config.Load().Repositories {
			if rclone.Check(repository) {
				rclone.Sync(repository)
			}
		}

		<-exit
		fmt.Println()
		return nil
	},
}

func init() {
	daemonCmd.AddCommand(daemonInstall)
	daemonCmd.AddCommand(daemonRemove)
	daemonCmd.AddCommand(daemonStart)
	rootCmd.AddCommand(daemonCmd)
}
