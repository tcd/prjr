package cmd

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"github.com/tcd/prjr/internal/prjr"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:     "rm",
	Aliases: []string{"remove"},
	Short:   "Remove an existing project",
	Long:    `Remove an existing project`,
	Run: func(cmd *cobra.Command, args []string) {
		projects, err := prjr.GetLocalProjects()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		noconfirm, _ := cmd.Flags().GetBool("noconfirm")
		if noconfirm {
			rmFuncNoconfirm(projects)
		} else {
			rmFunc(projects)
		}
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)

	rmCmd.Flags().BoolP("noconfirm", "Y", false, "Bypass any and all confirmation messages.")
}

func rmFunc(pjs prjr.Projects) {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	shouldRemove := false
	prompt := &survey.Confirm{
		Message: fmt.Sprintf("Remove project %q?", cwd),
	}
	survey.AskOne(prompt, &shouldRemove)

	if shouldRemove {
		pjs.RemoveByRoot(cwd)
		err = pjs.Save()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Project %q removed\n", cwd)
		}
	}
}

func rmFuncNoconfirm(pjs prjr.Projects) {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	pjs.RemoveByRoot(cwd)
	err = pjs.Save()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Project %q removed\n", cwd)
	}
}
