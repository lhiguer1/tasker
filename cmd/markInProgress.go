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

// markInProgressCmd represents the markInProgress command
var markInProgressCmd = &cobra.Command{
	Use:   "mark-in-progress <index>",
	Short: "Mark a tasks as `in-progress`.",
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

		service.MarkInProgress(i)

		if err := service.SaveTasks(); err != nil {
			fmt.Println("Error saving tasks:", err)
			return
		}

	},
}

func init() {
	rootCmd.AddCommand(markInProgressCmd)

}
