package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cwd, _ = os.Getwd()

var rootCmd = &cobra.Command{
	Use:   "rcloud",
	Short: "rcloud is a backup real time backup system",
	CompletionOptions: cobra.CompletionOptions{
		HiddenDefaultCmd: true,
	},
	DisableFlagsInUseLine: true,
}

func init() {
	rootCmd.SetHelpCommand(&cobra.Command{
		Hidden: true,
	})
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
