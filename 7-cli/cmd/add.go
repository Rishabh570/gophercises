package cmd

import (
	"fmt"
	"strings"
	"task/db"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(AddCmd)
}

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new task to the list",
	Run: func(cmd *cobra.Command, args []string) {
		taskText := strings.Join(args, " ")

		taskId := db.AddTask(taskText)

		fmt.Println("✅ Task created, id:", taskId)
	},
}
