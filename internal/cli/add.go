/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// What a ridiculous name
func AddAddCmd(parent *cobra.Command) *cobra.Command {
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
