package cmd

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/anotherhope/rcloud/app/internal"
	"github.com/anotherhope/rcloud/app/message"
	"github.com/anotherhope/rcloud/app/socket"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Args:  cobra.ExactValidArgs(0),
	Use:   "status",
	Short: "Show status of synchronized folders",
	RunE: func(cmd *cobra.Command, args []string) error {

		var output = [][]string{}
		var max = []int{9, 7, 5, 6, 11}
		output = append(output, []string{"RCLOUD ID", "ENABLED", "STATUS", "SOURCE", "DESTINATION"})

		for _, repository := range internal.App.Repositories {
			if max[0] < len(repository.Name[0:12]) {
				max[0] = len(repository.Name[0:12])
			}

			status := repository.GetStatus()

			client := socket.Client()
			if client != nil {

				m := &message.Message{
					Request:  message.ReqStatus(repository.Name),
					Response: &message.Response{},
				}

				client.Send(m)
				client.Close()
				status = m.Response.ToString()
			}

			if max[2] < len(status) {
				max[2] = len(status)
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
					status,
					repository.Source,
					repository.Destination,
				},
			)
		}

		if _, err := os.Stat(env.SocketPath); errors.Is(err, os.ErrNotExist) {
			fmt.Println("SERVICE: OFF")
		} else {
			fmt.Println("SERVICE: ON")
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
