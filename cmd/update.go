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

// updateCmdCmd represents the updateCmd command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		err := RunUpdateCmd(args)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	},
}

func RunUpdateCmd(args []string) error {
	if len(args) < 2 {
		return errors.New("usage: update <id> <description>")
	}

	id, _ := strconv.Atoi(args[0])
	description := args[1]

	return task.UpdateTask(id, description)
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
