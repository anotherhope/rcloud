package cmd

import (
	"fmt"

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
			// TODO
			fmt.Println("start")
			return nil
		},
	}

	stop := &cobra.Command{
		Args:  cobra.MaximumNArgs(1),
		Use:   "stop",
		Short: "Stop real time synchronization",
		RunE: func(cmd *cobra.Command, args []string) error {
			// TODO
			fmt.Println("stop")
			return nil
		},
	}

	rootCmd.AddCommand(cmd)
	cmd.AddCommand(start)
	cmd.AddCommand(stop)
}
