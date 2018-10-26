package cmd

import (
	"github.com/spf13/cobra"
)


func init() {
	RunCmd.AddCommand(benchMarkCmd)
}

var benchMarkCmd = &cobra.Command{
	Use:   "benchmark",
	Short: "submit a benchmark job to netbrain system",
	Long:  `submit a benchmark job to netbrain system`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}