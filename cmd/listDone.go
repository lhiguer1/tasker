/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/lhiguer1/tasker/tasks"
	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Lists all tasks marked `done`.",
	Run: func(cmd *cobra.Command, args []string) {
		service := tasks.TaskService{Filename: tasksFile}
		// Load tasks from the file initially
		if err := service.LoadTasks(); err != nil {
			fmt.Println("Error loading tasks:", err)
			return
		}
		service.ListDoneTasks()
	},
}

func init() {
	listCmd.AddCommand(doneCmd)
}
