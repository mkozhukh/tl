package cmd

import (
	"fmt"
	"strconv"
	"tl/db"

	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status <list-id>",
	Short: "Show task counts by state for a list",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		listID, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("invalid list-id: %w", err)
		}
		s, err := db.GetListStatus(conn, listID)
		if err != nil {
			return err
		}
		return printJSON(s)
	},
}
