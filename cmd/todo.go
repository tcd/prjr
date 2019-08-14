package cmd

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"github.com/tcd/prjr/internal/prjr"
)

// todoCmd represents the todo command
var todoCmd = &cobra.Command{
	Use:   "todo",
	Short: "List TODO comments in a project",
	Long:  `List TODO comments in a project`,
	Run: func(cmd *cobra.Command, args []string) {
		projects, err := prjr.GetLocalProjects()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		todoFunc(projects)
	},
}

func init() {
	rootCmd.AddCommand(todoCmd)
}

func todoFunc(pjs prjr.Projects) {
	root := ""
	prompt := &survey.Select{
		Message: "Which project would you like to check on?",
		Options: pjs.Roots(),
	}
	survey.AskOne(prompt, &root)
	if root != "" {
		pj, ok := pjs.FindByRoot(root)
		if ok {
			todos := pj.TODOs()
			for _, todo := range todos {
				fmt.Println(todo)
			}
		}
	}
	os.Exit(0)
}
