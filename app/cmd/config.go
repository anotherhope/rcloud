package cmd

import (
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Enter an interactive configuration session.",
		RunE: func(cmd *cobra.Command, args []string) error {
			sub := exec.Command("rclone", append([]string{"config"}, args[:]...)...)
			sub.Stdout = os.Stdout
			sub.Stdin = os.Stdin
			sub.Stderr = os.Stderr
			return sub.Run()
		},
		DisableFlagsInUseLine: true,
	}

	rootCmd.AddCommand(cmd)
}
