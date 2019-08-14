package cli

import (
	"bytes"
	"fmt"
	"strings"
	"text/tabwriter"

	"github.com/tcd/prjr/internal/prjr"
	"github.com/tcd/prjr/internal/todo"
)

// TabProjects returns a formatted list of Projects.
func TabProjects(pjs prjr.Projects) string {
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

// TabTodos returns a formatted list of Todos.
func TabTodos(todos []todo.Todo) string {
	var buf bytes.Buffer
	writer := tabwriter.NewWriter(&buf, 0, 8, 1, '\t', tabwriter.AlignRight)
	fmt.Fprintln(writer, strings.Join([]string{"type", "content", "path"}, "\t"))
	fmt.Fprintln(writer, strings.Join([]string{"----", "-------", "----"}, "\t"))
	for _, t := range todos {
		fmt.Fprintln(writer, strings.Join([]string{t.Type, t.Content, t.File}, "\t"))
	}
	err := writer.Flush()
	if err != nil {
		fmt.Println(err)
	}
	return buf.String()
}
