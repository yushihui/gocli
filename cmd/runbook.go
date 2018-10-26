package cmd

import (
	"github.com/spf13/cobra"
)


func init() {
	ShowCmd.AddCommand(runbookCmd)
}

var runbookCmd = &cobra.Command{
	Use:   "runbook",
	Short: "show the runbook information",
	Long:  `show the runbook information`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}