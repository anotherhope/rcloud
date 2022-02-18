package cmd

import (
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Enter in interactive configuration session. (alias: rclone config)",
	RunE: func(cmd *cobra.Command, args []string) error {
		sub := exec.Command("rclone", append([]string{"config"}, args[:]...)...)
		sub.Stdout = os.Stdout
		sub.Stdin = os.Stdin
		sub.Stderr = os.Stderr
		return sub.Run()
	},
	DisableFlagsInUseLine: true,
}

func init() {
	rootCmd.AddCommand(configCmd)
}
