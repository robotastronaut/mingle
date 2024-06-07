package cli

// Get information about the runtime environment, including Mudlet data

import (
	"fmt"
	"strings"

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
			fmt.Println(summarizeEnv(mudletInstance))
			return nil
		},
	}

	parent.AddCommand(envCmd)

	return envCmd
}

func summarizeEnv(m mudlet.Mudlet) string {
	summary := strings.Builder{}
	summary.WriteString(header.Render("Mudlet Profiles (" + m.ConfigPath + ")"))
	summary.WriteString("\n")
	for _, profile := range m.Profiles {
		summary.WriteString(summaryLine(profile.Name, profile.Path))
		for _, pkg := range profile.Packages {
			summary.WriteString(summaryLine("  (pkg) "+pkg.Name(), pkg.ConfigPath()))
		}

	}

	return summary.String()
}
