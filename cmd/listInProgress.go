/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/lhiguer1/tasker/tasks"
	"github.com/spf13/cobra"
)

// inProgressCmd represents the inProgress command
var inProgressCmd = &cobra.Command{
	Use:   "in-progress",
	Short: "Lists all tasks marked `in-progress`.",
	Run: func(cmd *cobra.Command, args []string) {
		service := tasks.TaskService{Filename: tasksFile}
		// Load tasks from the file initially
		if err := service.LoadTasks(); err != nil {
			fmt.Println("Error loading tasks:", err)
			return
		}
		service.ListInProgressTasks()
	},
}

func init() {
	listCmd.AddCommand(inProgressCmd)
}
