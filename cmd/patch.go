package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(patchCmd)
}

var patchCmd = &cobra.Command{
	Use:   "patch",
	Short: "deploy a patch to netbrain system",
	Long:  `deploy a patch to netbrain system`,
	Example: `
# TODO
  netbrain patch webserver
# TODO
  netbrain patch fsc`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}