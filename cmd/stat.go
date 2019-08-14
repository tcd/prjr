package cmd

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey"
	"github.com/spf13/cobra"
	"github.com/tcd/prjr/internal/prjr"
)

// statCmd represents the stat command
var statCmd = &cobra.Command{
	Use:   "stat",
	Short: "Print Git status information for a project",
	Long:  `Print Git status information for a project`,
	Run: func(cmd *cobra.Command, args []string) {
		projects, err := prjr.GetLocalProjects()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		statFunc(projects)

	},
}

func init() {
	rootCmd.AddCommand(statCmd)
}

func statFunc(pjs prjr.Projects) {
	root := ""
	prompt := &survey.Select{
		Message: "Which project would you like to check on?",
		Options: pjs.Roots(),
	}
	survey.AskOne(prompt, &root)
	if root != "" {
		pj, ok := pjs.FindByRoot(root)
		if ok {
			status, err := pj.GitStatus()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(status)
		}
	}
	os.Exit(0)
}
