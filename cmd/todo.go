package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"github.com/tcd/prjr/internal/cli"
	"github.com/tcd/prjr/internal/prjr"
)

var todoCmd = &cobra.Command{
	Use:   "todo",
	Short: "List TODO comments in a project",
	Run: func(cmd *cobra.Command, args []string) {
		root, _ := cmd.Flags().GetString("root")
		if root != "" {
			todoFuncFlag(root)
		}
		todoFuncPrompt()
	},
}

func init() {
	rootCmd.AddCommand(todoCmd)
	todoCmd.Flags().String("root", "", "Path to the root of the Project to search for TODOs")
}

func todoFuncPrompt() {
	pjs, err := prjr.GetLocalProjects()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

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
			fmt.Println(cli.TabTodos(todos))
		}
	}
	os.Exit(0)
}

func todoFuncFlag(root string) {
	pjs, err := prjr.GetLocalProjects()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if !filepath.IsAbs(root) {
		fmt.Fprintln(os.Stderr, "Please provide an absolute path")
		os.Exit(1)
	}

	pj, ok := pjs.FindByRoot(root)
	if ok {
		todos := pj.TODOs()
		fmt.Println(cli.TabTodos(todos))
	}
	os.Exit(0)
}
