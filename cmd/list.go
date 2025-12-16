/*
Copyright © 2025 vohongquan.dev@gmail.com
*/
package cmd

import (
	"errors"
	"fmt"
	"strconv"
	"task-tracker/internal/storage"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list [id]",
	Short: "List tasks",
	Long: `Display all tasks or a specific task by ID.

Examples:
  task list
  task list 3
`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		tasks, err := storage.LoadTasks()
		if err != nil {
			return err
		}

		// Case: list all
		if len(args) == 0 {
			for _, task := range tasks {
				fmt.Printf("[%d] %s (status: %v)\n", task.Id, task.Description, task.Status)
			}
			return nil
		}

		// Case: list by id
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return errors.New("invalid id, id must be an integer")
		}

		for _, task := range tasks {
			if task.Id == id {
				fmt.Printf("[%d] %s (completed: %v)\n", task.Id, task.Description, task.Status)
				return nil
			}
		}

		return fmt.Errorf("task with id %d not found", id)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
