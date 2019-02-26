package node

import (
	"github.com/dcos/dcos-cli/api"
	"github.com/dcos/dcos-core-cli/pkg/metrics"
	"github.com/dcos/dcos-core-cli/pkg/pluginutil"
	"github.com/spf13/cobra"
)

func newCmdNodeMetricsQuery(ctx api.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "query <query>",
		Short: "Print query",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return metrics.NewClient(pluginutil.HTTPClient("")).Query(args[0])
		},
	}
	return cmd
}
