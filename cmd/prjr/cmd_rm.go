package main

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey"
	"github.com/tcd/prjr/internal/prjr"
)

func rmCmd(pjs prjr.Projects) {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
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
