/*
Copyright © 2024 Leonel Higuera <lhiguer1@asu.edu>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tasker",
	Short: "Task tracker CLI app to track your tasks and manage your to-do list.",
}

var tasksFile string

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&tasksFile, "file", "tasks.json", "Tasks file")
}
