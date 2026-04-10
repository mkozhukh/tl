package cmd

import (
	"encoding/json"
	"fmt"
	"strconv"
	"github.com/mkozhukh/tl/db"

	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done <task-id>",
	Short: "Mark a task as done",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		taskID, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("invalid task-id: %w", err)
		}
		result, _ := cmd.Flags().GetString("result")
		if result != "" && !json.Valid([]byte(result)) {
			return fmt.Errorf("--result must be valid JSON")
		}
		task, err := db.CompleteTask(conn, taskID, result)
		if err != nil {
			return err
		}
		return printTask(task)
	},
}

func init() {
	doneCmd.Flags().String("result", "", "JSON result data")
}
