package main

import (
	"fmt"

	"github.com/tcd/prjr/internal/prjr"
)

func addCmd(pjs prjr.Projects) {
	pj, err := prjr.NewProjectHere()
	if err != nil {
		fmt.Println(err)
	}

	pjs.Add(pj)
	err = pjs.Save()
	if err != nil {
		fmt.Println(err)
	}
}
