/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/lhiguer1/tasker/tasks"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete <index>",
	Short: "Delete a task.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		service := tasks.TaskService{Filename: tasksFile}

		// Load tasks from the file initially
		if err := service.LoadTasks(); err != nil {
			fmt.Println("Error loading tasks:", err)
			return
		}
		i, err := strconv.Atoi(args[0])

		if err != nil {
			fmt.Println("Invalid index:", err)
			return
		}

		service.DeleteTask(i)

		if err := service.SaveTasks(); err != nil {
			fmt.Println("Error saving tasks:", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
