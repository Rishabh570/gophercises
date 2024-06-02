package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var encodingKey2 string

func init() {
	AddCmd.Flags().StringVarP(&encodingKey2, "encodingKey2", "k", "", "Encoding key to use")
	RootCmd.AddCommand(AddCmd)
}

var AddCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets a new secret",
	Run: func(cmd *cobra.Command, args []string) {
		taskText := strings.Join(args, " ")
		fmt.Println("taskText:", taskText)
		fmt.Println("Encoding key", encodingKey2)

		// taskId := db.AddTask(taskText)

		// fmt.Println("✅ Secret set")
	},
}
