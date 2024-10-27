package cmd

import (
	"cli-to-do-list/logic"

	"github.com/spf13/cobra"
)

var completeToDo = &cobra.Command{
	Use:   "complete",
	Short: "Mark a to-do as complete.",
	Long:  "Mark a to-do as complete to remove it from the list.",

	Run: func(cmd *cobra.Command, args []string) {
		logic.RemoveRecord()
	},
}

func init() {
	// Register the add command as a subcommand of the root command
	rootCmd.AddCommand(completeToDo)
}
