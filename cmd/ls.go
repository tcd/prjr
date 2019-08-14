package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tcd/prjr/internal/cli"
	"github.com/tcd/prjr/internal/prjr"
)

var lsCmd = &cobra.Command{
	Use:     "ls [OPTIONS]",
	Aliases: []string{"list"},
	Short:   "List all projects",
	Run: func(cmd *cobra.Command, args []string) {
		projects, err := prjr.GetLocalProjects()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if len(projects.P) == 0 {
			fmt.Println("No projects")
			os.Exit(0)
		}

		table, _ := cmd.Flags().GetBool("table")
		if table {
			listFuncTable(projects)
		}

		outFields, _ := cmd.Flags().GetStringSlice("fields")
		if len(outFields) > 0 {
			listFuncFields(projects, outFields)
		}

		listFunc(projects)
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
	lsCmd.Flags().StringSlice("fields", []string{}, "Fields to print for each project.")
	lsCmd.Flags().Bool("table", false, "Print projects in table format.")
}

func listFunc(pjs prjr.Projects) {
	if len(pjs.P) > 0 {
		fmt.Print(cli.TabProjects(pjs))
	} else {
		fmt.Println("No projects")
	}
	os.Exit(0)
}

func listFuncFields(pjs prjr.Projects, fields []string) {
	if len(pjs.P) > 0 {
		for _, pj := range pjs.P {
			var toPrint []string
			if contains(fields, "name") {
				toPrint = append(toPrint, pj.Name)
			}
			if contains(fields, "root") {
				toPrint = append(toPrint, pj.Root)
			}
			output := strings.Join(toPrint, "\t")
			fmt.Println(output)
		}
	} else {
		fmt.Println("No projects")
	}
	os.Exit(0)
}

func listFuncTable(pjs prjr.Projects) {
	if len(pjs.P) > 0 {
		cli.Table(pjs)
		os.Exit(0)
	} else {
		fmt.Println("No projects")
		os.Exit(0)
	}
}
