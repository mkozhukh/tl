package cmd

import (
	"fmt"
	"strconv"
	"tl/db"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get <task-id>",
	Short: "Get a single task",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		taskID, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("invalid task-id: %w", err)
		}
		task, err := db.GetTask(conn, taskID)
		if err != nil {
			return err
		}
		return printTask(task)
	},
}
