package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var cwd, _ = os.Getwd()

var rootCmd = &cobra.Command{
	Use:   "rcloud",
	Short: "Rcloud is a backup real time backup system",
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
		os.Exit(1)
	}
}