package cmd

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"github.com/tcd/prjr/internal/prjr"
)

var addCmd = &cobra.Command{
	Use:   "add [OPTIONS]",
	Short: "Add a new Project",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		projects, err := prjr.GetLocalProjects()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		cwd, _ := os.Getwd()
		_, ok := projects.FindByRoot(cwd)
		if ok {
			fmt.Println("This directory is already a project.")
			os.Exit(1)
		}

		noconfirm, _ := cmd.Flags().GetBool("noconfirm")
		if noconfirm {
			addFuncNoconfirm(projects)
		} else {
			addFunc(projects)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().BoolP("noconfirm", "Y", false, "Bypass any and all confirmation messages")
}

func addFunc(pjs prjr.Projects) {
	shouldAdd := false
	prompt := &survey.Confirm{Message: "Add the current directory as a project?"}
	survey.AskOne(prompt, &shouldAdd)

	if shouldAdd {
		pj, err := prjr.NewProjectHere()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		pjs.Add(pj)
		err = pjs.Save()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {
			fmt.Printf("Added new project %q\n", pj.Root)
			os.Exit(0)
		}
	}
	os.Exit(1)
}

func addFuncNoconfirm(pjs prjr.Projects) {
	pj, err := prjr.NewProjectHere()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pjs.Add(pj)
	err = pjs.Save()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Printf("Added new project %q\n", pj.Root)
		os.Exit(0)
	}
}
