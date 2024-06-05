/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"golang.org/x/mod/semver"
)

func AddInitCmd(parent *cobra.Command) *cobra.Command {
	modFile := Mfile{}
	initCmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize a new module",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			// Adding the 'v' here because this semver package expects it
			if !semver.IsValid("v" + modFile.Version) {
				return fmt.Errorf("mfile version is not a valid semver")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			modFile.Form().Run()
		},
	}

	parent.AddCommand(initCmd)

	initCmd.Flags().StringVarP(&modFile.Package, "package", "p", "", "package name")
	initCmd.Flags().StringVarP(&modFile.Version, "version", "v", "0.0.1", "package version")
	initCmd.Flags().StringVarP(&modFile.Description, "description", "d", "", "package description")
	initCmd.Flags().StringVarP(&modFile.Author, "author", "a", "", "package author")
	initCmd.Flags().StringVarP(&modFile.Title, "title", "t", "", "package title")
	initCmd.Flags().BoolVarP(&modFile.OutputFile, "outputfile", "o", true, "TODO outputfile")

	return initCmd
}
