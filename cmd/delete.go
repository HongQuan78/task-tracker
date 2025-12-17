/*
Copyright © 2025 vohongquan.dev@gmail.com
*/
package cmd

import (
	"fmt"
	"strconv"
	"task-tracker/internal/storage"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a task",
	Long: `Delete a task from the task tracker by its ID.

The specified task will be permanently removed and can no longer
be listed or updated.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("id is required; please provide an Id to delete the task")
		}
		id := args[0]
		if id == "" {
			fmt.Println("The Id is required. Please provide an Id to delete the task.")
		}
		deleteId, err := strconv.Atoi(id)
		if err != nil {
			return err
		}

		if err := storage.DeleteTask(deleteId); err != nil {
			return err
		}

		fmt.Println("✅ Task deleted.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
