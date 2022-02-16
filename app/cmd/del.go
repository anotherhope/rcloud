package cmd

import (
	"github.com/anotherhope/rcloud/app/internal"
	"github.com/spf13/cobra"
)

var delCmd = &cobra.Command{
	Args:  cobra.ExactValidArgs(1),
	Use:   "del <folder>",
	Short: "Delete synchronized folder",
	RunE: func(cmd *cobra.Command, args []string) error {
		return internal.Del(args[0])
	},
	DisableFlagsInUseLine: true,
}

func init() {
	rootCmd.AddCommand(delCmd)
}
