package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "secret",
	Short: "Secrets CLI",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("CLI to manage your secrets")
		fmt.Println()

		// flag.Parse()

		// Load dependencies
		// db.LoadDatabase()

		fmt.Println(cmd.UsageString())
	},
}
