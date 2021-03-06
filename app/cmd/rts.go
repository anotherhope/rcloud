package cmd

import (
	"fmt"
	"strings"

	"github.com/anotherhope/rcloud/app/internal/config"
	"github.com/anotherhope/rcloud/app/internal/repositories"
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
		for k, repository := range repositories.Repositories {
			if !repository.RTS {
				if len(args) == 1 && strings.HasPrefix(repository.Name, args[0]) || len(args) == 0 {
					count++
					repository.Start()
				}
			}

			repositories.Repositories[k] = repository
		}

		config.App.Set("repositories", repositories.Repositories)
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
		for k, repository := range repositories.Repositories {
			if repository.RTS {
				if len(args) == 1 && strings.HasPrefix(repository.Name, args[0]) || len(args) == 0 {
					count++
					repository.Stop()
				}
			}

			repositories.Repositories[k] = repository
		}

		config.App.Set("repositories", repositories.Repositories)
		fmt.Printf("%v synchronization(s) has been stopped\n", count)

		return nil
	},
}

func init() {
	rtsCmd.AddCommand(rtsStart)
	rtsCmd.AddCommand(rtsStop)
	rootCmd.AddCommand(rtsCmd)
}
