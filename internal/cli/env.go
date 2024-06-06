package cli

// Get information about the runtime environment, including Mudlet data

import (
	"fmt"

	"github.com/robotastronaut/mingle/internal/mudlet"
	"github.com/spf13/cobra"
)

func AddEnvCmd(parent *cobra.Command) *cobra.Command {
	envCmd := &cobra.Command{
		Use:   "env",
		Short: "Get environment details",

		RunE: func(cmd *cobra.Command, args []string) error {
			mudletInstance, err := mudlet.GetMudlet()
			if err != nil {
				return err
			}
			fmt.Println(mudletInstance)
			return nil
		},
	}

	parent.AddCommand(envCmd)

	return envCmd
}
