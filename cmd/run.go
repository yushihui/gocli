package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(RunCmd)
}

var RunCmd = &cobra.Command{
	Use:   "run",
	Short: "run a benchmark/path/discovery/runbook...",
	Long:  `run a benchmark/path/discovery/runbook...`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}
