package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "CLI task manager",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("task is a CLI for managing your TODOs.")
		fmt.Println()

		// Load dependencies
		// db.LoadDatabase()

		fmt.Println(cmd.UsageString())
	},
}
