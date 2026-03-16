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

// markDoneCmd represents the markDone command
var markDoneCmd = &cobra.Command{
	Use:   "mark-done",
	Short: "Mark a task as done",
	Run: func(cmd *cobra.Command, args []string) {
		err := RunMarkDoneCmd(args)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	},
}

func RunMarkDoneCmd(args []string) error {
	if len(args) < 1 {
		return errors.New("usage: mark-done <id>")
	}

	id, _ := strconv.Atoi(args[0])

	return task.UpdateTask(id, "", "done")
}

func init() {
	rootCmd.AddCommand(markDoneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// markDoneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// markDoneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
