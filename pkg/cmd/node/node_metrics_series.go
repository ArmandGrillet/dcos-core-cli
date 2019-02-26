package node

import (
	"github.com/dcos/dcos-cli/api"
	"github.com/dcos/dcos-core-cli/pkg/metrics"
	"github.com/dcos/dcos-core-cli/pkg/pluginutil"
	"github.com/spf13/cobra"
)

func newCmdNodeMetricsSeries(ctx api.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "series",
		Short: "Print series",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := metrics.NewClient(pluginutil.HTTPClient("")).Series()
			if err != nil {
				return err
			}
			return nil
		},
	}
	return cmd
}
