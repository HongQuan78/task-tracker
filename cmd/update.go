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

var description string

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update [id]",
	Short: "Update a task description by id",
	Long: `Update the description of a task by its id.
Pass the id as the argument and the new description via the flag.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]

		if id == "" || description == "" {
			return errors.New("id and description are required")
		}

		updateId, err := strconv.Atoi(id)

		if err != nil {
			return err
		}

		err = storage.UpdateTask(updateId, description)

		if err != nil {
			return err
		}

		fmt.Println("✅ Task Updated.")
		return nil

	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringVarP(&description, "description", "d", "", "Description")
}
