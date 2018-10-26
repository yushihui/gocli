package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(configCmd)
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Modify IE configs",
	Long:  `Modify IE configs`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}