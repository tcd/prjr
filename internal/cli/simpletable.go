package cli

import (
	"fmt"

	"github.com/alexeyco/simpletable"
	"github.com/tcd/prjr/internal/prjr"
)

// Table prints a pretty table of all projects.
// ⎇ ✓ ✔ ✗
// https://github.com/alexeyco/simpletable
func Table(pjs prjr.Projects) {
	var data = tableFormat(pjs)

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "Name"},
			{Align: simpletable.AlignCenter, Text: "Root"},
			{Align: simpletable.AlignCenter, Text: "Git Status"},
			{Align: simpletable.AlignCenter, Text: "TODOs"},
		},
	}

	for _, row := range data {
		r := []*simpletable.Cell{
			{Text: row[0].(string)},
			{Text: row[1].(string)},
			{Text: row[2].(string)},
			{Align: simpletable.AlignLeft, Text: fmt.Sprintf("%d", row[3].(int))},
		}
		table.Body.Cells = append(table.Body.Cells, r)
	}

	table.SetStyle(simpletable.StyleUnicode)
	table.Println()
}

func tableFormat(pjs prjr.Projects) [][]interface{} {
	var data [][]interface{}

	for _, pj := range pjs.P {
		var pjData = []interface{}{
			pj.Name,
			pj.Root,
		}
		if pj.VCS {
			gs, _ := pj.GitStatus()
			pjData = append(pjData, []interface{}{gs.String()}...)
		} else {
			pjData = append(pjData, []interface{}{""}...)
		}
		pjData = append(pjData, []interface{}{pj.TODOCount()}...)
		data = append(data, pjData)
	}
	return data
}
