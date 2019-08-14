package tui

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
			{Align: simpletable.AlignCenter, Text: "TODOs"},
			// {Align: simpletable.AlignCenter, Text: "Git Status"},
		},
	}

	for _, row := range data {
		r := []*simpletable.Cell{
			{Text: row[0].(string)},
			{Text: row[1].(string)},
			{Align: simpletable.AlignLeft, Text: fmt.Sprintf("%d", row[2].(int))},
			// {Text: row[3].(string)},
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
		pjData = append(pjData, []interface{}{pj.TODOCount()}...)
		// if pj.VCS {
		// 	pjData = append(pjData, []interface{}{"✔"}...)
		// } else {
		// 	pjData = append(pjData, []interface{}{""}...)
		// }
		data = append(data, pjData)
	}
	return data
}
