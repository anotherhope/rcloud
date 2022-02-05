package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{
		Use:   "daemon",
		Short: "Daemon management",
	}

	install := &cobra.Command{
		Args:  cobra.ExactArgs(0),
		Use:   "install",
		Short: "Install daemon",
		RunE: func(cmd *cobra.Command, args []string) error {
			// TODO
			fmt.Println("install")
			return nil
		},
	}

	remove := &cobra.Command{
		Args:  cobra.ExactArgs(0),
		Use:   "remove",
		Short: "Remove daemon",
		RunE: func(cmd *cobra.Command, args []string) error {
			// TODO
			fmt.Println("remove")
			return nil
		},
	}

	standalone := &cobra.Command{
		Args:  cobra.ExactArgs(0),
		Use:   "standalone",
		Short: "Run daemon in standalone mode",
		RunE: func(cmd *cobra.Command, args []string) error {
			// TODO
			fmt.Println("standalone")
			return nil
		},
	}

	rootCmd.AddCommand(cmd)
	cmd.AddCommand(install)
	cmd.AddCommand(remove)
	cmd.AddCommand(standalone)
}
