package cmd

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/anotherhope/rcloud/app/internal"
	"github.com/spf13/cobra"
)

var directoryCmd = &cobra.Command{
	Use:   "directory",
	Short: "Manage directory",
}

var directoryStart = &cobra.Command{
	Args:  cobra.ExactValidArgs(2),
	Use:   "add <source> <destination>",
	Short: "Add synchronized directory",
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

		return internal.Add(&internal.Directory{
			Name:        fmt.Sprintf("%x", h.Sum(nil)),
			Source:      source,
			Destination: destination,
			Watch:       100 * time.Millisecond,
		})
	},
	DisableFlagsInUseLine: true,
}

var directoryStop = &cobra.Command{
	Args:  cobra.ExactValidArgs(1),
	Use:   "del <directory>",
	Short: "Delete synchronized directory",
	RunE: func(cmd *cobra.Command, args []string) error {
		return internal.Del(args[0])
	},
	DisableFlagsInUseLine: true,
}

func init() {
	directoryCmd.AddCommand(directoryStart)
	directoryCmd.AddCommand(directoryStop)
	rootCmd.AddCommand(directoryCmd)
}
