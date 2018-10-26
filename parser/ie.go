package parser

import (
	"bytes"
	"encoding/json"
	"github.com/buger/jsonparser"
	"github.com/olekukonko/tablewriter"
	"os"
)

func Ie2Json(response []byte) (string, error) {
	content, _, _, _ := jsonparser.Get(response, "data")
	var out bytes.Buffer
	err := json.Indent(&out, content, "", "    ")
	return string(out.Bytes()), err

}

func TableOutPut(data [][]string, header []string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)

	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}

func PathOutput(path Path) string {

	var buffer bytes.Buffer
	for i, hop := range path.Hops {
		buffer.WriteString("\t")
		buffer.WriteString(hop.Name)
		if i < len(path.Hops)-1 {
			buffer.WriteString("\n\t    |- ")
			buffer.WriteString(hop.South)
			if !path.Success && i == len(path.Hops)-2 {
				buffer.WriteString("\n\t    |\n\t    \x1b[0;31mX\x1b[0m- ")
			} else {
				buffer.WriteString("\n\t    |\n\t    |- ")
			}
			buffer.WriteString(path.Hops[i+1].North)
			buffer.WriteString("\n")
		}
	}
	return buffer.String()
}
