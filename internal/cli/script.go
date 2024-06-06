/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func ScriptCmd(parent *cobra.Command) *cobra.Command {
	scriptCmd := &cobra.Command{
		Use:   "script",
		Short: "Add a new script",

		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("script called")
		},
	}

	parent.AddCommand(scriptCmd)

	return scriptCmd
}
