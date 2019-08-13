package main

import (
	"fmt"

	"github.com/AlecAivazis/survey"
	"github.com/tcd/prjr/internal/prjr"
)

func todoCmd(pjs prjr.Projects) {
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
}
