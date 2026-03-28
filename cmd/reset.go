package cmd

import (
	"fmt"
	"strconv"
	"tl/db"

	"github.com/spf13/cobra"
)

var resetCmd = &cobra.Command{
	Use:   "reset <task-id>",
	Short: "Reset a failed or active task to pending",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		taskID, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("invalid task-id: %w", err)
		}
		task, err := db.ResetTask(conn, taskID)
		if err != nil {
			return err
		}
		return printJSON(task)
	},
}
