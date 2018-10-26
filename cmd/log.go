package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(logCmd)
}

var logCmd = &cobra.Command{
	Use:   "log",
	Short: "print the logs of a component",
	Long:  `describe the resources`,
	Example: `
# Return snapshot logs from webserver
  netbrain logs webserver
# export fsc log
  netbrain logs fsc -o fsc.log`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}
