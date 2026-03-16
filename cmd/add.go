/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"task-tracker/internal/task"

	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Create an new task",
	Run: func(cmd *cobra.Command, args []string) {
		err := RunAddCmd(args)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	},
}

func RunAddCmd(args []string) error {
	if len(args) < 1 {
		return errors.New("task name is required")
	}

	description := args[0]
	return task.AddTask(description)
}

func init() {
	rootCmd.AddCommand(AddCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// AddCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// AddCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
