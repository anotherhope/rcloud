package rts

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Stop = &cobra.Command{
	Args:  cobra.ExactArgs(1),
	Use:   "stop",
	Short: "Stop real time synchronization",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("stop")
		return nil
	},
}
