package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(status)
}

var status = &cobra.Command{
	Args:  cobra.ExactValidArgs(2),
	Use:   "status",
	Short: "Show status of synchronized folders",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	DisableFlagsInUseLine: true,
}
