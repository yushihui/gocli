package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(delCmd)
}

var delCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete the resource from netbrain system",
	Long:  `delete the resource from netbrain system`,
	Example: `
# delete a map by name
  netbrain delete map --name=m1
# delete a map by path
  netbrain delete map --path=p1`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}