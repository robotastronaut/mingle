/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cli

import (
	"os"

	"github.com/spf13/cobra"
)

func AddInitCmd(parent *cobra.Command) *cobra.Command {
	gen := Generator{}
	var nonInteractive bool
	initCmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize a new module",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			// Adding the 'v' here because this semver package expects it
			// if !semver.IsValid("v" + modFile.Version) {
			// 	return fmt.Errorf("version is not a valid semver")
			// }

			// if len(modFile.Package) < 1 {
			// 	return fmt.Errorf("invalid package name")
			// }

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			workdir, err := os.Getwd()
			if err != nil {
				return err
			}

			if len(args) > 0 && len(args[0]) > 0 {
				workdir = args[0]
			}

			gen.Path = workdir

			if !nonInteractive {
				gen.Form().Run()
				gen.Confirm().Run()

				if !gen._ready {
					return nil
				}
			}

			err = gen.Run()
			if err != nil {
				return err
			}
			// TODO: Check for cleanup

			return nil
		},
	}

	parent.AddCommand(initCmd)

	initCmd.Flags().StringVarP(&gen.Name, "package", "p", "", "Package name")
	initCmd.Flags().StringVarP(&gen.Description, "description", "d", "", "Package description")
	initCmd.Flags().StringVarP(&gen.Author, "author", "a", "", "Package author")
	initCmd.Flags().StringVarP(&gen.Title, "title", "t", "", "Package title")
	initCmd.Flags().BoolVarP(&gen.OutputFile, "watchable", "w", true, "Enable creation of Muddle watch file")
	initCmd.Flags().BoolVarP(&nonInteractive, "noninteractive", "n", false, "Run in non-interactive mode")

	return initCmd
}
