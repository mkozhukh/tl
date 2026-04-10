package cmd

import (
	"fmt"
	"strconv"
	"tl/db"

	"github.com/spf13/cobra"
)

var failCmd = &cobra.Command{
	Use:   "fail <task-id>",
	Short: "Mark a task as failed",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		taskID, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("invalid task-id: %w", err)
		}
		reason, _ := cmd.Flags().GetString("reason")
		task, err := db.FailTask(conn, taskID, reason)
		if err != nil {
			return err
		}
		return printTask(task)
	},
}

func init() {
	failCmd.Flags().String("reason", "", "Reason for failure")
}
