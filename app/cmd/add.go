package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

var cwd, _ = os.Getwd()

func init() {
	rootCmd.AddCommand(add)
}

var add = &cobra.Command{
	Args:  cobra.ExactValidArgs(2),
	Use:   "add <source> <destination>",
	Short: "Add to synchronized folder (" + cwd + ")",
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
		fmt.Println("rclone", "sync", args[0], args[1], "--dry-run")
		//output, err := exec.Command("rclone", "sync", args[0], args[1], "--dry-run").Output()
		command := exec.Command("rclone", "sync", args[0], args[1], "--dry-run", "--progress")
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr

		return command.Run()
	},
	DisableFlagsInUseLine: true,
}
