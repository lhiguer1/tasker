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

// markDoneCmd represents the markDone command
var markDoneCmd = &cobra.Command{
	Use:   "mark-done <index>",
	Short: "Mark a task as `done`.",
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

		service.MarkDone(i)

		if err := service.SaveTasks(); err != nil {
			fmt.Println("Error saving tasks:", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(markDoneCmd)
}
