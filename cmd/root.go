/*
Copyright © 2025 vohongquan.dev@gmail.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "A simple CLI to track your tasks",
	Long: `Task Tracker is a lightweight command-line application
that helps you manage your tasks directly from the terminal.

You can add, list, update, mark tasks as done, and delete tasks.
All tasks are stored locally in a JSON file, making the tool fast,
portable, and easy to use without any external dependencies.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
