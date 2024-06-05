/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var modPath string

// Root returns the base command when called without any subcommands
// TODO: Validate path
func Root() *cobra.Command {
	// rootCmd is the core Cobra command struct
	rootCmd := &cobra.Command{
		Use:   "muddler-go",
		Short: "Go implementation of demonnic/muddler",
		Args:  cobra.MaximumNArgs(1),
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			_, err := os.Stat(modPath)
			if err != nil {
				return fmt.Errorf("invalid module path")
			}
			return nil
		},
		// Uncomment the following line if your bare application
		// has an action associated with it:
		RunE: func(cmd *cobra.Command, args []string) error {
			// Get config and print
			c, err := LoadMfile(modPath)
			if err != nil {
				return err
			}

			fmt.Println(c.String())

			return nil
		},
	}

	rootCmd.PersistentFlags().StringVarP(&modPath, "directory", "D", ".", "root directory of module")
	AddInitCmd(rootCmd)
	return rootCmd
}

/**

With no flags, will print info about module
Final parameter regardless of subcommand is the directory to run in, default to "."
With `generate` will enter TUI
With `-n <name>` creates a new module named `<name>`
Additional params:
	alias <name REQUIRED>
		-a <active>, boolean, default true
		-c <command>, string, command to send to mud
		-r <regex>, string, regular expression. Slashes escaped
		-s <script>, string, lua script to run. If not provided, looks for <alias.name>.lua



*/
