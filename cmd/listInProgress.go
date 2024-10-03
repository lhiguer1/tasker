/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// inProgressCmd represents the inProgress command
var inProgressCmd = &cobra.Command{
	Use:   "in-progress",
	Short: "Lists all tasks marked `in-progress`.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("inProgress called")
	},
}

func init() {
	listCmd.AddCommand(inProgressCmd)
}
