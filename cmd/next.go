package cmd

import (
	"fmt"
	"strconv"
	"tl/db"

	"github.com/spf13/cobra"
)

var nextCmd = &cobra.Command{
	Use:   "next <list-id> [task-id]",
	Short: "Claim the next pending task, or a specific task by ID",
	Args:  cobra.RangeArgs(1, 2),
	RunE: func(cmd *cobra.Command, args []string) error {
		owner, _ := cmd.Flags().GetString("owner")
		var task *db.Task
		var err error
		if len(args) == 2 {
			taskID, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return fmt.Errorf("invalid task-id: %w", err)
			}
			task, err = db.ClaimTask(conn, taskID, owner)
		} else {
			listID, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("invalid list-id: %w", err)
			}
			task, err = db.ClaimNextTask(conn, listID, owner)
		}
		if err != nil {
			return err
		}
		return printTask(task)
	},
}

func init() {
	nextCmd.Flags().String("owner", "", "ID of the agent claiming the task")
}
