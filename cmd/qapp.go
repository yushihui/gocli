package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yushihui/cli/parser"
)

func init() {
	initRun()
	initGet()
}

func initRun() {
	var name, target string
	var qappRCmd = &cobra.Command{
		Use:   "qapp",
		Short: "submit a qapp job to netbrain system",
		Long:  `submit a qapp job to netbrain system`,
		Run: func(cmd *cobra.Command, args []string) {
			q, _ := cmd.Flags().GetString("name")
			d, _ := cmd.Flags().GetString("target")

			fmt.Printf("Qapp %s is running against: %s ...\n", q, d)
		},
	}
	qappRCmd.Flags().StringVarP(&name, "name", "n", "", "qapp name")
	qappRCmd.Flags().StringVarP(&target, "target", "t", "", "target devices")
	RunCmd.AddCommand(qappRCmd)
}

func initGet() {

	var query string
	var qappGCmd = &cobra.Command{
		Use:   "qapp",
		Short: "display all qapps",
		Long:  `display all qapps`,
		Run: func(cmd *cobra.Command, args []string) {
			header,data:=getQapps()
			parser.TableOutPut(data, header)
		},
	}
	qappGCmd.Flags().StringVarP(&query, "query", "q", "", "qapp filter")

	GetCmd.AddCommand(qappGCmd)
}

func getQapps()([]string, [][]string){
	header:=[]string{"Name","Type","Version"}
	data := [][]string{
		[]string{"overall-health-monitor", "Monitor", "7.1"},
		[]string{"collect-ospf-log", "data", "7.1"},
		[]string{"collect-Qos", "data", "3.6"},
		[]string{"monitor-cpu", "Monitor", "3.7"},
		[]string{"monitor-memory", "Monitor", "5.0"},
		[]string{"collect-vlan-info", "data", "8.0"},
		[]string{"collect-vrf-cisco", "data", "7.1"},
	}
	return header, data
}