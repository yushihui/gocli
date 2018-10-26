package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yushihui/cli/client"
	"github.com/yushihui/cli/parser"
)

func init() {
	var device string
	var deviceCmd = &cobra.Command{
		Use:   "device",
		Short: "get device information",
		Long:  `get device information`,
		Run: func(cmd *cobra.Command, args []string) {
			id:=cmd.Flag("ip").Value.String()
			//a217428c-9944-430c-8c1b-8d4ce9a53971
			res := client.Post("DataModel/getSimpleDevicesByGuids", []string{id})
			fmt.Println(parser.Ie2Json(res.Body()))

		},
	}
	deviceCmd.Flags().StringVarP(&device, "ip", "i", "9", "please input device ip.")
	GetCmd.AddCommand(deviceCmd)
}
