/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cli

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/robotastronaut/mpm/internal/muddler"
	"github.com/spf13/cobra"
)

func Root() *cobra.Command {
	// rootCmd is the core Cobra command struct
	rootCmd := &cobra.Command{
		Use:   "mpm",
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
			module, err := muddler.FindModule(workdir)

			if err != nil {
				return err
			} else {
				fmt.Println(summarizeModule(module))
			}
			fmt.Println("\nSCRIPTS: \n", module.Scripts)
			return nil
		},
	}

	AddInitCmd(rootCmd)
	AddEnvCmd(rootCmd)
	AddScriptCmd(rootCmd)
	return rootCmd
}

func summarizeModule(m *muddler.Module) string {
	summary := strings.Builder{}
	summary.WriteString(header.Render("Package Summary"))
	summary.WriteString("\n")
	summary.WriteString(summaryLine("Path", m.Path))
	summary.WriteString(summaryLine("Name", m.Package))
	summary.WriteString(summaryLine("Title", m.Title))
	summary.WriteString(summaryLine("Author", m.Author))
	summary.WriteString(summaryLine("OutputFile", strconv.FormatBool(m.OutputFile)))
	desc := m.Description
	if desc == "" {
		desc = "<empty>"
	}
	summary.WriteString(summaryBlock("Description", desc))

	return summary.String()
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
