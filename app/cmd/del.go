package cmd

import (
	"github.com/anotherhope/rcloud/app/interfaces"
	"github.com/spf13/cobra"
)

var delCmd = &cobra.Command{
	Args:  cobra.ExactValidArgs(1),
	Use:   "del <folder>",
	Short: "Delete synchronized folder",
	RunE: func(cmd *cobra.Command, args []string) error {
		return interfaces.Del(args[0])
	},
	DisableFlagsInUseLine: true,
}

func init() {
	rootCmd.AddCommand(delCmd)
}
