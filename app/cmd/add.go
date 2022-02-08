package cmd

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/anotherhope/rcloud/app/config"
	"github.com/anotherhope/rcloud/app/repositories"
	"github.com/spf13/cobra"
)

var add_cmd = &cobra.Command{
	Args:  cobra.ExactValidArgs(2),
	Use:   "add <source> <destination>",
	Short: "Add to synchronized folder (" + cwd + ")",
	RunE: func(cmd *cobra.Command, args []string) error {

		var err error
		var source = args[0]
		var destination = args[1]

		if source, err = repositories.IsValid(source); err != nil {
			return err
		}

		if destination, err = repositories.IsValid(destination); err != nil {
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

		//fmt.Println("rclone", "sync", source, destination, "--dry-run")
		//output, err := exec.Command("rclone", "sync", args[0], args[1], "--dry-run").Output()
		//command := exec.Command("rclone", "sync", source, destination, "--dry-run", "--progress")
		//command.Stdout = os.Stdout
		//command.Stderr = os.Stderr

		//return //command.Run()
		//return nil
	},
	DisableFlagsInUseLine: true,
}

func init() {
	rootCmd.AddCommand(add_cmd)
}
