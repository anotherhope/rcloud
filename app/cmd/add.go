package cmd

import (
	"crypto/sha1"
	"fmt"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/anotherhope/rcloud/app/config"
	"github.com/anotherhope/rcloud/app/repositories"
	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{
		Args:  cobra.ExactValidArgs(2),
		Use:   "add <source> <destination>",
		Short: "Add to synchronized folder (" + cwd + ")",
		RunE: func(cmd *cobra.Command, args []string) error {

			var source = args[0]
			var destination = args[1]

			if strings.Contains(source, ":") {
				remote := strings.Split(source, ":")[0]
				output, err := exec.Command("rclone", "listremotes").Output()
				if err != nil {
					return err
				}

				availableChoices := strings.Split(string(output), "\n")
				if i := sort.SearchStrings(availableChoices, remote+":"); i == 0 {
					return fmt.Errorf("rclone remote not available for source")
				}
			} else {
				source, _ = filepath.Abs(args[0])
			}

			if strings.Contains(destination, ":") {
				remote := strings.Split(destination, ":")[0]
				output, err := exec.Command("rclone", "listremotes").Output()
				if err != nil {
					return err
				}

				availableChoices := strings.Split(string(output), "\n")
				if i := sort.SearchStrings(availableChoices, remote+":"); i > 0 {
					return fmt.Errorf("rclone remote not available for destination")
				}
			} else {
				destination, _ = filepath.Abs(args[1])
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

	rootCmd.AddCommand(cmd)
}
