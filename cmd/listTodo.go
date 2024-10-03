/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/lhiguer1/tasker/tasks"
	"github.com/spf13/cobra"
)

// todoCmd represents the todo command
var todoCmd = &cobra.Command{
	Use:   "todo",
	Short: "Lists all tasks marked `todo`.",
	Run: func(cmd *cobra.Command, args []string) {
		service := tasks.TaskService{Filename: tasksFile}
		// Load tasks from the file initially
		if err := service.LoadTasks(); err != nil {
			fmt.Println("Error loading tasks:", err)
			return
		}
		service.ListTodoTasks()
	},
}

func init() {
	listCmd.AddCommand(todoCmd)
}
