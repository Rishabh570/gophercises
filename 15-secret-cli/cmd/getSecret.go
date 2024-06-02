package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var encodingKey string

func init() {
	GetCmd.Flags().StringVarP(&encodingKey, "encodingKey", "k", "", "Encoding key to use")

	RootCmd.AddCommand(GetCmd)
}

var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Fetches a secret",
	Run: func(cmd *cobra.Command, args []string) {
		taskText := strings.Join(args, " ")
		fmt.Println("taskText:", taskText)
		fmt.Println("Encoding key", encodingKey)

		// taskId := db.AddTask(taskText)

		// fmt.Println("✅ Read secret")
	},
}
