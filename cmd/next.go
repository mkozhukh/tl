package cmd

import (
	"fmt"
	"strconv"
	"tl/db"

	"github.com/spf13/cobra"
)

var nextCmd = &cobra.Command{
	Use:   "next <list-id>",
	Short: "Claim the next pending task",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		listID, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("invalid list-id: %w", err)
		}
		task, err := db.ClaimNextTask(conn, listID)
		if err != nil {
			return err
		}
		return printJSON(task)
	},
}
