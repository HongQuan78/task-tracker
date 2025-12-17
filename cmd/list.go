/*
Copyright © 2025 vohongquan.dev@gmail.com
*/
package cmd

import (
	"fmt"
	"task-tracker/internal/model"
	"task-tracker/internal/storage"

	"github.com/spf13/cobra"
)

var statusFilter string

const timeFmt = "2-1-2006 15:04:05"

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List tasks",
	Long: `Display all tasks.

Examples:
  task list
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		tasks, err := storage.LoadTasks()
		if err != nil {
			return err
		}

		filter := model.TaskStatus("")

		if statusFilter != "" {
			switch statusFilter {
			case string(model.StatusTodo), string(model.StatusInProgress), string(model.StatusDone):
				filter = model.TaskStatus(statusFilter)
			default:
				return fmt.Errorf("invalid status, use: todo | in-progress | done")
			}
		}

		if len(args) == 0 {
			for _, task := range tasks {
				if filter != "" && task.Status != filter {
					continue
				}
				fmt.Printf(
					"[%d] %s (status: %v) \ncreated: %s \nupdated: %s\n",
					task.Id,
					task.Description,
					task.Status,
					task.CreatedAt.Format(timeFmt),
					task.UpdatedAt.Format(timeFmt),
				)
			}
			return nil
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringVarP(&statusFilter, "status", "s", "", "Filter by status: todo|in-progress|done")
}
