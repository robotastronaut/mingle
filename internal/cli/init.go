/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cli

import (
	"os"

	"github.com/spf13/cobra"
)

func AddInitCmd(parent *cobra.Command) *cobra.Command {
	// modFile := NewMFile()
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

			gen := Generator{
				Path: workdir,
			}

			gen.Form().Run()

			err = gen.Run()
			if err != nil {
				return err
			}
			// TODO: Check for cleanup

			return nil
		},
	}

	parent.AddCommand(initCmd)

	// initCmd.Flags().StringVarP(&modFile.Package, "package", "p", "", "package name")
	// initCmd.MarkFlagRequired("package")
	// modFile._viper.BindPFlag("package", initCmd.Flags().Lookup("package"))

	// initCmd.Flags().StringVarP(&modFile.Version, "version", "v", "0.0.1", "package version")
	// modFile._viper.BindPFlag("version", initCmd.Flags().Lookup("version"))

	// initCmd.Flags().StringVarP(&modFile.Description, "description", "d", "", "package description")
	// modFile._viper.BindPFlag("description", initCmd.Flags().Lookup("description"))

	// initCmd.Flags().StringVarP(&modFile.Author, "author", "a", "", "package author")
	// modFile._viper.BindPFlag("author", initCmd.Flags().Lookup("author"))

	// initCmd.Flags().StringVarP(&modFile.Title, "title", "t", "", "package title")
	// modFile._viper.BindPFlag("title", initCmd.Flags().Lookup("title"))

	return initCmd
}
