package main

import (
	"fmt"
	"os"
	"task/cmd"
	"task/db"
)

func main() {
	// Initialize database
	dbInstance, err := db.LoadDatabase()
	if err != nil {
		fmt.Println("Could not initialize DB")
		os.Exit(1)
	}

	// Close the connection when process exits
	defer dbInstance.Close()

	// Run root CLI command
	cmd.RootCmd.Execute()
}
