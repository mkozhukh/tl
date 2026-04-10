package cmd

import (
	"encoding/json"
	"fmt"
	"strconv"
	"tl/db"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add <list-id> <title>",
	Short: "Add a task to a list",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		listID, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("invalid list-id: %w", err)
		}
		meta, _ := cmd.Flags().GetString("meta")
		if meta != "" && !json.Valid([]byte(meta)) {
			return fmt.Errorf("--meta must be valid JSON")
		}
		task, err := db.AddTask(conn, listID, args[1], meta)
		if err != nil {
			return err
		}
		return printJSON(map[string]int64{"id": task.ID})
	},
}

func init() {
	addCmd.Flags().String("meta", "", "JSON metadata for the task")
}
