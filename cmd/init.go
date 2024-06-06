/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"golang.org/x/mod/semver"
)

func AddInitCmd(parent *cobra.Command) *cobra.Command {
	modFile := NewMFile()
	initCmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize a new module",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			// Adding the 'v' here because this semver package expects it
			if !semver.IsValid("v" + modFile.Version) {
				return fmt.Errorf("version is not a valid semver")
			}

			if len(modFile.Package) < 1 {
				return fmt.Errorf("invalid package name")
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			// whatever
			p := tea.NewProgram(initialTUIModel())

			if _, err := p.Run(); err != nil {
				log.Fatal(err)
			}

			// runPrompt()

			// Immediately use viper to save the file
			return modFile._viper.WriteConfigAs("mfile")
		},
	}

	parent.AddCommand(initCmd)

	initCmd.Flags().StringVarP(&modFile.Package, "package", "p", "", "package name")
	initCmd.MarkFlagRequired("package")
	modFile._viper.BindPFlag("package", initCmd.Flags().Lookup("package"))

	initCmd.Flags().StringVarP(&modFile.Version, "version", "v", "0.0.1", "package version")
	modFile._viper.BindPFlag("version", initCmd.Flags().Lookup("version"))

	initCmd.Flags().StringVarP(&modFile.Description, "description", "d", "", "package description")
	modFile._viper.BindPFlag("description", initCmd.Flags().Lookup("description"))

	initCmd.Flags().StringVarP(&modFile.Author, "author", "a", "", "package author")
	modFile._viper.BindPFlag("author", initCmd.Flags().Lookup("author"))

	initCmd.Flags().StringVarP(&modFile.Title, "title", "t", "", "package title")
	modFile._viper.BindPFlag("title", initCmd.Flags().Lookup("title"))

	return initCmd
}
