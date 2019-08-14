package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tcd/prjr/internal/prjr"
	"github.com/tcd/prjr/internal/tui"
)

func listCmd(pjs prjr.Projects) {
	listCommand := flag.NewFlagSet("list", flag.ExitOnError)
	longFlag := listCommand.Bool("l", false, "Provide a longer listing for projects.")
	// fieldFlag := listCommand.String("fields", "root", "Project fields to list")

	listCommand.Parse(os.Args[2:])

	if *longFlag == true {
		fmt.Println("TODO: long list. Accept args for fields to list?")
	} else {
		if len(pjs.P) > 0 {
			// for _, project := range pjs.P {
			// 	fmt.Println(project.Root)
			// }
			tui.Table(pjs)
		} else {
			fmt.Println("No projects")
		}
	}
}
