package rts

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Start = &cobra.Command{
	Args:  cobra.ExactArgs(1),
	Use:   "start",
	Short: "Start real time synchronization",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("start")
		return nil
	},
}
