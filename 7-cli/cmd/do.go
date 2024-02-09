package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"task/db"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(DoCmd)
}

var DoCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task as completed",
	Run: func(cmd *cobra.Command, args []string) {
		var taskIdsToMarkAsCompleted []int
		for _, item := range args {
			taskId, err := strconv.Atoi(item)
			if err != nil {
				log.Fatal("could not convert string to int")
				os.Exit(1)
			}
			taskIdsToMarkAsCompleted = append(taskIdsToMarkAsCompleted, taskId)
		}

		for _, id := range taskIdsToMarkAsCompleted {
			err := db.DoTask(id)
			if err != nil {
				fmt.Println("Could not check off the task")
				os.Exit(1)
			}
			fmt.Println("✅ Task", id, "is marked as completed")
		}
	},
}
