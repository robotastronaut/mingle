/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func AddCmd(parent *cobra.Command) *cobra.Command {
	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Add a new feature",

		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("add called")
		},
	}

	parent.AddCommand(addCmd)

	return addCmd
}
