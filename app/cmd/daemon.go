package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var daemon_cmd = &cobra.Command{
	Use:   "daemon",
	Short: "Daemon management",
}

var daemon_install = &cobra.Command{
	Args:  cobra.ExactArgs(0),
	Use:   "install",
	Short: "Install daemon",
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO
		fmt.Println("install")
		return nil
	},
}

var daemon_remove = &cobra.Command{
	Args:  cobra.ExactArgs(0),
	Use:   "remove",
	Short: "Remove daemon",
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO
		fmt.Println("remove")
		return nil
	},
}

var daemon_standalone = &cobra.Command{
	Args:  cobra.ExactArgs(0),
	Use:   "standalone",
	Short: "Run daemon in standalone mode",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("standalone")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(daemon_cmd)
	daemon_cmd.AddCommand(daemon_install)
	daemon_cmd.AddCommand(daemon_remove)
	daemon_cmd.AddCommand(daemon_standalone)
}
