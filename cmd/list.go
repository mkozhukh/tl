package cmd

import (
	"tl/db"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all task lists with counts",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		lists, err := db.GetAllLists(conn)
		if err != nil {
			return err
		}
		return printJSON(lists)
	},
}
