/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"strconv"
	"task-tracker/internal/storage"

	"github.com/spf13/cobra"
)

var status string

var statusCmd = &cobra.Command{
	Use:   "status [id]",
	Short: "Update a task status by id",
	Long: `Update the status of a task by its id.
Valid statuses: todo, in-progress, done.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]
		if status == "" || id == "" {
			return errors.New("status and id is required")
		}

		updateId, err := strconv.Atoi(id)
		if err != nil {
			return err
		}

		if err := storage.ChangeStatus(status, updateId); err != nil {
			return err
		}

		fmt.Println("✅ Status updated.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
	statusCmd.Flags().StringVar(&status, "set-status", "", "Update status: todo|in-progress|done")

}
