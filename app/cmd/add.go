package cmd

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/anotherhope/rcloud/app/config"
	"github.com/anotherhope/rcloud/app/repositories"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Args:  cobra.ExactValidArgs(2),
	Use:   "add <source> <destination>",
	Short: "Add to synchronized folder (" + cwd + ")",
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		var source = args[0]
		var destination = args[1]

		if source, err = repositories.IsValid(source, false); err != nil {
			return err
		}

		if destination, err = repositories.IsValid(destination, true); err != nil {
			return err
		}

		h := sha1.New()
		h.Write([]byte(source + destination))

		return repositories.Add(&config.Directory{
			Name:        fmt.Sprintf("%x", h.Sum(nil)),
			Source:      source,
			Destination: destination,
			Watch:       100 * time.Millisecond,
		})
	},
	DisableFlagsInUseLine: true,
}

func init() {
	rootCmd.AddCommand(addCmd)
}
