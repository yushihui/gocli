package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(loginCmd)
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login to netbrain",
	Long:  `login to netbrain`,
	Example: `
# login with user name & pwd
  netbrain login --user admin -p admin`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}
