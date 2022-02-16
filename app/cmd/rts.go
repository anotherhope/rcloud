package cmd

import (
	"fmt"
	"strings"

	"github.com/anotherhope/rcloud/app/interfaces"
	"github.com/spf13/cobra"
)

var rtsCmd = &cobra.Command{
	Use:   "rts",
	Short: "Real Time Synchronization",
}

var rtsStart = &cobra.Command{
	Args:  cobra.MaximumNArgs(1),
	Use:   "start",
	Short: "Start real time synchronization",
	RunE: func(cmd *cobra.Command, args []string) error {
		count := 0
		for k, repository := range interfaces.App.Repositories {
			if !repository.RTS {
				if len(args) == 1 && strings.HasPrefix(repository.Name, args[0]) || len(args) == 0 {
					count++
					if err := repository.Start(); err != nil {
						return err
					}
				}
			}

			interfaces.App.Repositories[k] = repository
		}

		interfaces.App.Set("repositories", interfaces.App.Repositories)
		fmt.Printf("%v synchronization(s) has been started\n", count)

		return nil
	},
}

var rtsStop = &cobra.Command{
	Args:  cobra.MaximumNArgs(1),
	Use:   "stop",
	Short: "Stop real time synchronization",
	RunE: func(cmd *cobra.Command, args []string) error {
		count := 0
		for k, repository := range interfaces.App.Repositories {
			if repository.RTS {
				if len(args) == 1 && strings.HasPrefix(repository.Name, args[0]) || len(args) == 0 {
					count++
					if err := repository.Stop(); err != nil {
						return err
					}
				}
			}

			interfaces.App.Repositories[k] = repository
		}

		interfaces.App.Set("repositories", interfaces.App.Repositories)
		fmt.Printf("%v synchronization(s) has been stopped\n", count)

		return nil
	},
}

func init() {
	rtsCmd.AddCommand(rtsStart)
	rtsCmd.AddCommand(rtsStop)
	rootCmd.AddCommand(rtsCmd)
}
