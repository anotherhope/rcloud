package cmd

import (
	"github.com/anotherhope/rcloud/app/repositories"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(del)
}

var del = &cobra.Command{
	Args:  cobra.ExactValidArgs(1),
	Use:   "del <folder>",
	Short: "Delete synchronized folder",
	RunE: func(cmd *cobra.Command, args []string) error {
		return repositories.Del(args[0])
	},
	DisableFlagsInUseLine: true,
}
