package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(ShowCmd)
}

var ShowCmd = &cobra.Command{
	Use:   "show",
	Short: "describe the resource",
	Long:  `describe the resource`,
	Example: `
# show device detail pane
  netbrain show device --ip 10.101.10.1
# show map info
  netbrain show map --name map1`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}
