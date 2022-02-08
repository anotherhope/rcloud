package cmd

import (
	"fmt"

	"github.com/anotherhope/rcloud/app/config"
	"github.com/anotherhope/rcloud/app/repositories"
	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{
		Use:   "rts",
		Short: "Real Time Synchronization",
	}

	start := &cobra.Command{
		Args:  cobra.MaximumNArgs(1),
		Use:   "start",
		Short: "Start real time synchronization",
		RunE: func(cmd *cobra.Command, args []string) error {
			repo := repositories.List()
			count := 0
			for k, repository := range repo {
				if !repository.RTS {
					if len(args) == 1 && repository.Name == args[0] || len(args) == 0 {
						count++
						if err := repository.Start(); err != nil {
							return err
						}
					}
				}

				repo[k] = repository
			}

			config.Set("repositories", repo)

			if count > 0 {
				fmt.Printf("%v synchronization(s) are started\n", count)
			}

			return nil
		},
	}

	stop := &cobra.Command{
		Args:  cobra.MaximumNArgs(1),
		Use:   "stop",
		Short: "Stop real time synchronization",
		RunE: func(cmd *cobra.Command, args []string) error {
			repo := repositories.List()
			count := 0
			for k, repository := range repositories.List() {
				if repository.RTS {
					if len(args) == 1 && repository.Name == args[0] || len(args) == 0 {
						count++
						if err := repository.Stop(); err != nil {
							return err
						}
					}
				}

				repo[k] = repository
			}

			config.Set("repositories", repo)

			if count > 0 {
				fmt.Printf("%v synchronization(s) are stopped\n", count)
			}

			return nil
		},
	}

	rootCmd.AddCommand(cmd)
	cmd.AddCommand(start)
	cmd.AddCommand(stop)
}
