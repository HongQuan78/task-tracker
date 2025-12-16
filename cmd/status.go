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

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status [id]",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
