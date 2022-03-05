package cmd

import (
	"os"
	"os/signal"

	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Args:  cobra.ExactArgs(0),
	Use:   "test",
	Short: "Test part of code",
	RunE: func(cmd *cobra.Command, args []string) error {
		var exit = make(chan os.Signal, 1)
		signal.Notify(exit, os.Interrupt)
		// algo a tester
		<-exit
		return nil
	},
	DisableFlagsInUseLine: true,
}

func init() {
	rootCmd.AddCommand(testCmd)
}
