package cli

import (
	"bytes"
	"fmt"
	"strings"
	"text/tabwriter"

	"github.com/tcd/prjr/internal/prjr"
)

// Tab returns a formatted list of Projects.
func Tab(pjs prjr.Projects) string {
	var buf bytes.Buffer
	writer := tabwriter.NewWriter(&buf, 0, 8, 1, '\t', tabwriter.AlignRight)
	for _, pj := range pjs.P {
		fmt.Fprintln(writer, strings.Join([]string{pj.Name, pj.Root}, "\t"))
	}
	err := writer.Flush()
	if err != nil {
		fmt.Println(err)
	}
	return buf.String()
}
