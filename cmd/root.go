package cmd

import (
	"database/sql"
	"errors"
	"os"
	"github.com/mkozhukh/tl/db"

	"github.com/spf13/cobra"
)

var conn *sql.DB

var rootCmd = &cobra.Command{
	Use:           "tl",
	Short:         "Task list manager for AI agents",
	SilenceErrors: true,
	SilenceUsage:  true,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		path := os.Getenv("TL_DB")
		if path == "" {
			path = "tl.db"
		}
		var err error
		conn, err = db.Open(path)
		return err
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		if conn != nil {
			conn.Close()
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(nextCmd)
	rootCmd.AddCommand(doneCmd)
	rootCmd.AddCommand(failCmd)
	rootCmd.AddCommand(statusCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(resetCmd)
	rootCmd.AddCommand(tasksCmd)
}

func Execute() int {
	if err := rootCmd.Execute(); err != nil {
		if errors.Is(err, db.ErrNoTasks) {
			printError(err)
			return 2
		}
		printError(err)
		return 1
	}
	return 0
}
