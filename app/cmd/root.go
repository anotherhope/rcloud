package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rcloud",
	Short: "Rcloud is a real time backup system",
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

// Execute is the entrypoint of all available commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
