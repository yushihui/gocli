package cmd

import (
	"github.com/spf13/cobra"
)


func init() {
	ShowCmd.AddCommand(mapCmd)
}

var mapCmd = &cobra.Command{
	Use:   "map",
	Short: "show the map information",
	Long:  `show the map information`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}