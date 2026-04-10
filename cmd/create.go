package cmd

import (
	"tl/db"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create <name>",
	Short: "Create a new task list",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		list, err := db.CreateList(conn, args[0])
		if err != nil {
			return err
		}
		return printJSON(map[string]int64{"id": list.ID})
	},
}
