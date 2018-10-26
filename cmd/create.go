package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a resource(map, runbook, qapp, dataview)",
	Long:  `Create a resource`,
	Example: `
# create a map from map.json
  netbrain create -f map.json
# create a runbook from runbook.yaml
  netbrain create -f runbook.yaml`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}