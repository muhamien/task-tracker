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

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		err := RunDeleteCmd(args)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	},
}

func RunDeleteCmd(args []string) error {
	if len(args) < 1 {
		return errors.New("usage: delete <id>")
	}

	id, _ := strconv.Atoi(args[0])

	return task.DeleteTask(id)
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
