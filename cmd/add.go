/*
Copyright © 2025 vohongquan.dev@gmail.com
*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [title]",
	Short: "Add a new task",
	Long: `Add a new task to the task tracker.

The task will be stored locally and can be listed, updated,
marked as done, or deleted later.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		title := args[0]

		if title == "" {
			return errors.New("task title cannot be empty")
		}

		fmt.Println("✅ Task added:", title)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
