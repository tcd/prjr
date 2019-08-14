package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tcd/prjr/internal/prjr"
	"github.com/tcd/prjr/internal/tui"
)

var lsCmd = &cobra.Command{
	Use:     "ls",
	Aliases: []string{"list"},
	Short:   "List all projects",
	Long:    `List all projects`,
	Run: func(cmd *cobra.Command, args []string) {
		// Set up
		projects, err := prjr.GetLocalProjects()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// Can't list projects if none exist.
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
		fmt.Print(tui.Tab(pjs))
	} else {
		fmt.Println("No projects")
	}
	os.Exit(0)
}

func listFuncFields(pjs prjr.Projects, fields []string) {
	if len(pjs.P) > 0 {
		for _, pj := range pjs.P {
			if contains(fields, "name") {
				fmt.Print(pj.Name)
				fmt.Print("\t")
			}
			if contains(fields, "root") {
				fmt.Print(pj.Root)
				fmt.Print("\t")
			}
			fmt.Print("\n")
		}
	} else {
		fmt.Println("No projects")
	}
	os.Exit(0)
}

func listFuncTable(pjs prjr.Projects) {
	if len(pjs.P) > 0 {
		tui.Table(pjs)
	} else {
		fmt.Println("No projects")
	}
	os.Exit(0)
}
