package cmd

import (
	"cli-to-do-list/logic"

	"github.com/spf13/cobra"
)

// Run the todo app in the CLI with "go run main.go todo".
var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "Todo is a simple CLI to-do list.",
	Long:  "A CLI to-do list that helps you keep track of your day to day tasks straight from the command line.",

	Run: func(cmd *cobra.Command, args []string) {
		// Start TUI
		logic.ListRecords()
	},
}

// Execute runs the root command
func Execute() error {
	return rootCmd.Execute()
}
