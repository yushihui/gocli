package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yushihui/cli/parser"
)


func init() {
	ShowCmd.AddCommand(ieCmd)
}

var ieCmd = &cobra.Command{
	Use:   "ie",
	Short: "show ie deployment",
	Long:  `show ie deployment`,
	Run: func(cmd *cobra.Command, args []string) {
		header,data:=getCluster()
		parser.TableOutPut(data, header)

	},
}

// return header & content
func getCluster()([]string, [][]string){
	header:=[]string{"Host","Role","Version", "Status","IP"}
	data := [][]string{
		[]string{"windows1", "WebServer", "7.1", "Running", "104.207.208.79"},
		[]string{"windows1", "Worker", "7.1", "Running", "104.207.208.79"},
		[]string{"Linux1", "Mongo", "3.6", "Running", "104.207.208.78"},
		[]string{"Linux1", "RabbitMQ", "3.7", "Running", "104.207.208.78"},
		[]string{"Linux1", "Redis", "5.0", "Running", "104.207.208.78"},
		[]string{"Linux1", "Task Engine", "8.0", "Running", "104.207.208.78"},
		[]string{"Linux1", "FSC", "7.1", "Running", "104.207.208.78"},
	}
	return header, data
}