package cmd

import (
	"cli-to-do-list/tui"

	"github.com/spf13/cobra"
)

var addToDo = &cobra.Command{
	Use:   "add",
	Short: "Add a to-do to the list.",
	Long:  "Add a to-do item to your list of things to get done.",

	Run: func(cmd *cobra.Command, args []string) {
		tui.AddRecord()
	},
}

func init() {
	// Register the add command as a subcommand of the root command
	rootCmd.AddCommand(addToDo)
}
