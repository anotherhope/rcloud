package cmd

import (
	"fmt"

	"github.com/anotherhope/rcloud/app/repositories"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(status)
}

var status = &cobra.Command{
	Args:  cobra.ExactValidArgs(0),
	Use:   "status",
	Short: "Show status of synchronized folders",
	RunE: func(cmd *cobra.Command, args []string) error {
		var output = [][]string{}
		var max = []int{4, 5, 6, 11}
		output = append(output, []string{"NAME", "STATUS", "SOURCE", "DESTINATION"})
		for _, repository := range repositories.List() {
			if max[0] < len(repository.Name) {
				max[0] = len(repository.Name)
			}

			if max[1] < len(repository.Status()) {
				max[1] = len(repository.Status())
			}

			if max[2] < len(repository.Source) {
				max[2] = len(repository.Source)
			}

			if max[3] < len(repository.Destination) {
				max[3] = len(repository.Destination)
			}

			output = append(output, []string{repository.Name, repository.Status(), repository.Source, repository.Destination})
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
