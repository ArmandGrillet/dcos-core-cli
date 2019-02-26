package node

import (
	"fmt"
	"strings"

	"github.com/dcos/dcos-cli/api"
	"github.com/dcos/dcos-core-cli/pkg/metrics"
	"github.com/dcos/dcos-core-cli/pkg/pluginutil"
	"github.com/spf13/cobra"
)

func newCmdNodeMetricsValues(ctx api.Context) *cobra.Command {
	var filter string
	var jsonOutput bool
	cmd := &cobra.Command{
		Use:   "values",
		Short: "Print values",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			values, err := metrics.NewClient(pluginutil.HTTPClient("")).Values()
			if err != nil {
				return err
			}
			for _, value := range values.Data {
				if strings.Contains(value, filter) {
					fmt.Println(value)
				}
			}
			return nil
		},
	}
	cmd.Flags().StringVar(&filter, "filter", "", "Filter the values")
	cmd.Flags().BoolVar(&jsonOutput, "json", false, "Print in json format")
	return cmd
}
