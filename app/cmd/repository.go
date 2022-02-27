package cmd

import (
	"crypto/sha1"
	"fmt"

	"github.com/anotherhope/rcloud/app/internal"
	"github.com/spf13/cobra"
)

var repositoryCmd = &cobra.Command{
	Use:   "repository",
	Short: "Manage repository",
}

var repositoryStart = &cobra.Command{
	Args:  cobra.ExactValidArgs(2),
	Use:   "add <source> <destination>",
	Short: "Add synchronized repository",
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		var source = args[0]
		var destination = args[1]

		if source, err = internal.IsValid(source, false); err != nil {
			return err
		}

		if destination, err = internal.IsValid(destination, true); err != nil {
			return err
		}

		h := sha1.New()
		h.Write([]byte(source + destination))

		return internal.Add(&internal.Repository{
			Name:        fmt.Sprintf("%x", h.Sum(nil)),
			Source:      source,
			Destination: destination,
		})
	},
	DisableFlagsInUseLine: true,
}

var repositoryStop = &cobra.Command{
	Args:  cobra.ExactValidArgs(1),
	Use:   "del <repository>",
	Short: "Delete synchronized repository",
	RunE: func(cmd *cobra.Command, args []string) error {
		return internal.Del(args[0])
	},
	DisableFlagsInUseLine: true,
}

func init() {
	repositoryCmd.AddCommand(repositoryStart)
	repositoryCmd.AddCommand(repositoryStop)
	rootCmd.AddCommand(repositoryCmd)
}
