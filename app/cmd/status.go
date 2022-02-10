package cmd

import (
	"fmt"
	"strconv"

	"github.com/anotherhope/rcloud/app/config"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Args:  cobra.ExactValidArgs(0),
	Use:   "status",
	Short: "Show status of synchronized folders",
	RunE: func(cmd *cobra.Command, args []string) error {
		var output = [][]string{}
		var max = []int{9, 6, 5, 6, 11}
		output = append(output, []string{"RCLOUD ID", "ACTIVE", "STATUS", "SOURCE", "DESTINATION"})
		for _, repository := range config.Load().Repositories {
			if max[0] < len(repository.Name[0:12]) {
				max[0] = len(repository.Name[0:12])
			}

			if max[2] < len(repository.Status()) {
				max[2] = len(repository.Status())
			}

			if max[3] < len(repository.Source) {
				max[3] = len(repository.Source)
			}

			if max[4] < len(repository.Destination) {
				max[4] = len(repository.Destination)
			}

			output = append(
				output,
				[]string{
					repository.Name[0:12],
					strconv.FormatBool(repository.RTS),
					repository.Status(),
					repository.Source,
					repository.Destination,
				},
			)
		}

		for _, line := range output {
			for i, col := range line {
				fmt.Printf("%-"+fmt.Sprint(max[i]+3)+"s", col)
			}
			fmt.Println()
		}

		return nil
	},
	DisableFlagsInUseLine: true,
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
