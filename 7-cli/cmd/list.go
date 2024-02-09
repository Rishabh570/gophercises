package cmd

import (
	"fmt"
	"os"
	"task/db"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(ListCmd)
}

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all the open tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.ListTasks()
		if err != nil {
			fmt.Println("Error while listing tasks: ", err)
			os.Exit(1)
		}

		if len(tasks) == 0 {
			fmt.Println("🎉 No tasks on your todo list")
			os.Exit(1)
		}

		fmt.Println("📃 Tasks:")
		for _, task := range tasks {
			fmt.Println("id: ", task.Id, " task: ", task.Text)
		}

	},
}
