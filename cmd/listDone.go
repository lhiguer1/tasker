/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Lists all tasks marked `done`.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("done called")
	},
}

func init() {
	listCmd.AddCommand(doneCmd)
}
