package debug

import (
	"github.com/dcos/dcos-cli/api"
	"github.com/dcos/dcos-core-cli/pkg/cmd/marathon/python"
	"github.com/spf13/cobra"
)

func newCmdMarathonDebugDetails(ctx api.Context) *cobra.Command {
	var json bool

	cmd := &cobra.Command{
		Use:   "details",
		Short: "Display detailed information for a queued instance launch for debugging purpose.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return python.InvokePythonCLI(ctx)
		},
	}

	cmd.Flags().BoolVar(&json, "json", false, "Print JSON-formatted data.")

	return cmd
}