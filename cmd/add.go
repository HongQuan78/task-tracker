/*
Copyright © 2025 vohongquan.dev@gmail.com
*/
package cmd

import (
	"errors"
	"fmt"
	"task-tracker/internal/storage"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [description]",
	Short: "Add a new task",
	Long: `Add a new task to the task tracker.

The task will be stored locally and can be listed, updated,
marked as done, or deleted later.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("id is required. Please provide an ID to delete the task")
		}
		description := args[0]

		if description == "" {
			return errors.New("task description cannot be empty")
		}
		err := storage.AddTask(description)
		if err != nil {
			return err
		}
		fmt.Println("✅ Task added:", description)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
