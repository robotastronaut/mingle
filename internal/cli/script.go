/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cli

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
)

func AddScriptCmd(parent *cobra.Command) *cobra.Command {
	scriptCmd := &cobra.Command{
		Use:   "script",
		Short: "Add a new script",

		RunE: func(cmd *cobra.Command, args []string) error {
			pattern := ""
			if len(args) > 0 {
				pattern = args[0]
			}
			manifestPaths, err := filepath.Glob(pattern)
			if err != nil {
				return err
			}
			fmt.Println(manifestPaths)
			return nil
		},
	}

	parent.AddCommand(scriptCmd)

	return scriptCmd
}
