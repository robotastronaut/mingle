/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func KeybindCmd(parent *cobra.Command) *cobra.Command {
	kbCmd := &cobra.Command{
		Use:   "keybind",
		Short: "Add a new keybind",

		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("keybind called")
		},
	}

	parent.AddCommand(kbCmd)

	return kbCmd
}
