/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cli

import (
	"fmt"
	"os"

	"github.com/robotastronaut/mingle/internal/mpackage"
	"github.com/spf13/cobra"
)

func Root() *cobra.Command {
	// rootCmd is the core Cobra command struct
	rootCmd := &cobra.Command{
		Use:   "mingle",
		Short: "Go implementation of demonnic/muddler",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			workdir, err := os.Getwd()
			if err != nil {
				return err
			}

			if len(args) > 0 && len(args[0]) > 0 {
				workdir = args[0]
			}

			// Get config and print
			module, err := mpackage.FindModule(workdir)

			if err != nil {
				return err
			} else {
				fmt.Println(module)
			}

			// fmt.Println(c.String())

			return nil
		},
	}

	AddInitCmd(rootCmd)
	AddEnvCmd(rootCmd)
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
