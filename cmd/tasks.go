package cmd

import (
	"fmt"
	"strconv"
	"tl/db"

	"github.com/spf13/cobra"
)

var tasksCmd = &cobra.Command{
	Use:   "tasks <list-id>",
	Short: "List all tasks in a list",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		listID, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("invalid list-id: %w", err)
		}
		status, _ := cmd.Flags().GetString("status")
		owner, _ := cmd.Flags().GetString("owner")
		tasks, err := db.GetListTasks(conn, listID, status, owner)
		if err != nil {
			return err
		}
		md, _ := cmd.Flags().GetBool("markdown")
		if md {
			for _, t := range tasks {
				fmt.Printf("- [%d] %s\n", t.ID, t.Title)
			}
			return nil
		}
		return printTasks(tasks)
	},
}

func init() {
	tasksCmd.Flags().Bool("markdown", false, "Output as markdown list")
	tasksCmd.Flags().String("status", "", "Filter by status (pending, active, done, failed)")
	tasksCmd.Flags().String("owner", "", "Filter by owner ID")
}
