package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yushihui/cli/parser"
)

func init() {
	var src, dst, protocol string
	var good string
	var pathCmd = &cobra.Command{
		Use:   "path",
		Short: "calculate the A-B path",
		Long:  `calculate the A-B path`,
		Example: `
# calculate path from dev1 to dev2
  netbrain run path -src dev1 -dst dev2
# calculate path from 10.10.10.1 to 10.10.100.1 via UDP
  netbrain run path -s 10.10.10.1 -d 10.10.100.1 -p UDP`,
		Run: func(cmd *cobra.Command, args []string) {
			src := cmd.Flag("src").Value.String()
			dst := cmd.Flag("dst").Value.String()
			gd, _ := cmd.Flags().GetString("good")
			hops := []parser.Hop{
				{Name: src, South: "e0", North: "e5"},
				{Name: "Bos-Dist-Nex5k-01", South: "e0", North: "e6"},
				{Name: "Bos-Core-6500", South: "e2", North: "e7"},
				{Name: "10.20.2.113", South: "e3", North: "e5"},
				{Name: dst, South: "e0", North: "e6"},
			}
			p := parser.Path{Hops: hops, Success: gd=="true"}
			fmt.Println(parser.PathOutput(p))
		},
	}
	pathCmd.Flags().StringVarP(&src, "src", "s", "", "source ip/host.")
	pathCmd.Flags().StringVarP(&dst, "dst", "d", "", "destination ip/host.")
	pathCmd.Flags().StringVarP(&protocol, "protocol", "p", "TCP", "protocol.")
	pathCmd.Flags().StringVarP(&good, "good", "g", "true", "good path(testing purpose)")
	RunCmd.AddCommand(pathCmd)
}
