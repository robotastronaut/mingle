/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package mpackage

import (
	"fmt"

	"github.com/spf13/cobra"
)

func TriggerCmd(parent *cobra.Command) *cobra.Command {
	triggerCmd := &cobra.Command{
		Use:   "trigger",
		Short: "Add a new trigger",

		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("trigger called")
		},
	}

	parent.AddCommand(triggerCmd)

	return triggerCmd
}
