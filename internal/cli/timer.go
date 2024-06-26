/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func TimerCmd(parent *cobra.Command) *cobra.Command {
	timerCmd := &cobra.Command{
		Use:   "timer",
		Short: "Add a new timer",

		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("timer called")
		},
	}

	parent.AddCommand(timerCmd)

	return timerCmd
}
