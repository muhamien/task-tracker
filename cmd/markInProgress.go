/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"strconv"
	"task-tracker/internal/task"

	"github.com/spf13/cobra"
)

// markInProgressCmd represents the markInProgress command
var markInProgressCmd = &cobra.Command{
	Use:   "mark-in-progress",
	Short: "Mark a task as in progress",
	Run: func(cmd *cobra.Command, args []string) {
		err := RunMarkInProgressCmd(args)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	},
}

func RunMarkInProgressCmd(args []string) error {
	if len(args) < 1 {
		return errors.New("usage: mark-in-progress <id>")
	}

	id, _ := strconv.Atoi(args[0])

	return task.UpdateTask(id, "", "in-progress")
}

func init() {
	rootCmd.AddCommand(markInProgressCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// markInProgressCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// markInProgressCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
