package cmd

import (
	"github.com/anotherhope/rcloud/app/cmd/rts"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(run)
	run.AddCommand(rts.Start)
	run.AddCommand(rts.Stop)
}

var run = &cobra.Command{
	Use:   "rts",
	Short: "Real Time Synchronization",
}
