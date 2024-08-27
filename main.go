package main

import (
	"fmt"
	"os"
)

func main() {
	// parse command line arguments passed in
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./main <task_id>")
		os.Exit(1)
	}
	// get the api key from the environment
	apiKey := os.Getenv("CLICKUP_API_KEY")
	if apiKey == "" {
		fmt.Println("Please set the CLICKUP_API_KEY environment variable")
		os.Exit(1)
	}
	taskId := os.Args[1]
	fmt.Printf("Task ID: %s\n", taskId)
}
