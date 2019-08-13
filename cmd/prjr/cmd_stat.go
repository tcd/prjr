package main

import (
	"fmt"

	"github.com/AlecAivazis/survey"
	"github.com/davecgh/go-spew/spew"
	"github.com/tcd/prjr/internal/prjr"
)

func statCmd(pjs prjr.Projects) {
	root := ""
	prompt := &survey.Select{
		Message: "Which project would you like a status for?",
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
			spew.Dump(status)
		}
	}
}
