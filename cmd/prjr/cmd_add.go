package main

import (
	"fmt"

	"github.com/AlecAivazis/survey"
	"github.com/tcd/prjr/internal/prjr"
)

func addCmd(pjs prjr.Projects) {
	shouldAdd := false
	prompt := &survey.Confirm{
		Message: "Add the current directory as a project?",
	}
	survey.AskOne(prompt, &shouldAdd)

	if shouldAdd {
		pj, err := prjr.NewProjectHere()
		if err != nil {
			fmt.Println(err)
		}

		pjs.Add(pj)
		err = pjs.Save()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Added new project %q\n", pj.Root)
		}
	}
}
