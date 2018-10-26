package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yushihui/cli/client"
	"github.com/yushihui/cli/parser"
)

func init() {
	RootCmd.AddCommand(GetCmd)
}

var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "get resources from netbrain system",
	Long:  `get resources from netbrain system`,
	Run: func(cmd *cobra.Command, args []string) {
		if args[0] == "overview" {
			//res:=client.Post("system/domain/domainsummary",`{"tenantId":"e6aea4c3-b950-6811-0f1b-381d029cfca5", "domainId":"69b322b9-540f-468c-9e28-f48cc60a6e48"}`)
		} else if args[0] == "tag" {
			res := client.Get("runbook-template/tags")
			fmt.Println(parser.Ie2Json(res.Body()))
			//fmt.Println(res.String())
		}

	},
}
