package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tcd/prjr/internal/prjr"
)

func listCmd(pjs prjr.Projects) {
	listCommand := flag.NewFlagSet("list", flag.ExitOnError)
	longFlag := listCommand.Bool("l", false, "Provide a longer listing for projects.")

	listCommand.Parse(os.Args[2:])

	if *longFlag == true {
		fmt.Println("TODO: long list. Accept args for fields to list?")
	} else {
		if len(pjs) > 0 {
			for _, project := range pjs {
				fmt.Println(project.Root)
			}
		} else {
			fmt.Println("No projects")
		}
	}
}
