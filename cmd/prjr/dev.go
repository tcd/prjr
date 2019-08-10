package main

import (
	"fmt"
	"os"

	"github.com/tcd/prjr/internal/prjr"
)

func test1() {
	pj := []prjr.Project{
		prjr.Project{
			Name: "prjr",
			Root: "/Users/clay/go/src/github.com/tcd/prjr/main.go",
			VCS:  true,
		},
	}

	err := prjr.SaveProjects(pj)
	if err != nil {
		fmt.Println(err)
	}
}

func test2() {
	projects, err := prjr.GetProjects()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(projects)
}

func test3() {
	var projects prjr.Projects

	existingProjects, err := prjr.GetProjects()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	projects.Add(existingProjects...)

	pj, err := prjr.NewProjectHere()
	if err != nil {
		fmt.Println(err)
	}

	projects.Add(pj)

	fmt.Println(projects)
}
