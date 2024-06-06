/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package mpackage

import (
	"fmt"

	"github.com/spf13/cobra"
)

// aliasCmd represents the alias command
func AliasCmd(parent *cobra.Command) *cobra.Command {
	aliasCmd := &cobra.Command{
		Use:   "alias",
		Short: "Add a new alias",

		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("alias called")
		},
	}

	parent.AddCommand(aliasCmd)

	return aliasCmd
}
