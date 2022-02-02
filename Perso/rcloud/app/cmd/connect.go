package cmd

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(connect)
}

var connect = &cobra.Command{
	Args:  cobra.ExactValidArgs(2),
	Use:   "connect <source> <destination>",
	Short: "Connecting to cloud",
	PreRunE: func(cmd *cobra.Command, args []string) error {

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
				return fmt.Errorf("rclone remote not availaible for source")
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
			fmt.Println(availableChoices, remote+":")
			if i := sort.SearchStrings(availableChoices, remote+":"); i > 0 {
				return fmt.Errorf("rclone remote not availaible for destination")
			}
		} else {
			destination, _ = filepath.Abs(args[1])
		}

		fmt.Println("PreRunE", source, destination)
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Run")
		output, err := exec.Command("rclone", "sync", args[0], args[1]).Output()
		fmt.Println(output)
		return err
	},
	DisableFlagsInUseLine: true,
}
