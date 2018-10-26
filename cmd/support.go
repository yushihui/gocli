package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(supportCmd)
}

var supportCmd = &cobra.Command{
	Use:   "support",
	Short: "get netbrain support information",
	Long:  `get netbrain support information`,
	Run: func(cmd *cobra.Command, args []string) {

		data := map[string]string{
			"Tel":     "+1 (800) 605-7964",
			"Email":   "support@netbraintech.com",
			"Website": "http://www.netbrain.com",
			"Fax":     "+1 (781) 998-5800",
			"Address": `15 Network Drive 2nd Floor Burlington, MA 01803 USA`,
		}
		content, error := json.MarshalIndent(data, "", "    ")
		if error == nil {
			fmt.Println(string(content))
		}

	},
}
