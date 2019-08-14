package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tcd/prjr/internal/prjr"
)

func init() {
	flag.Usage = fancyUsage
}

func main() {
	var projects prjr.Projects
	existingProjects, err := prjr.GetProjects()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	projects.Add(existingProjects...)

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "add":
			addCmd(projects)
		case "list", "ls":
			listCmd(projects)
		case "remove", "rm":
			rmCmd(projects)
		case "stat":
			statCmd(projects)
		case "todo":
			todoCmd(projects)
		default:
			usage()
		}
	} else {
		flag.Usage()
	}
}

func usage() {
	fmt.Printf("Usage: %s <command>\n", os.Args[0])
	fmt.Println("Commands: add, list, ls remove, rm, stat")
	flag.PrintDefaults()
}

func fancyUsage() {
	fmt.Println(titleString())
	fmt.Printf("Usage: %s <command>\n", os.Args[0])
	fmt.Println(`Commands:

add
	Add a new project
list, ls
	List existing projects
remove, rm
	Remove an existing project
stat
	Print Git info for a project
todo
	Search a project for TODO comments`)
	flag.PrintDefaults()
}

func titleString() string {
	return `   _______    _______        ___   _______
  |   __ "\  /"      \      |"  | /"      \
  (. |__) :)|:        |     ||  ||:        |
  |:  ____/ |_____/   )     |:  ||_____/   )
  (|  /      //      /   ___|  /  //      /
 /|__/ \    |:  __   \  /  :|_/ )|:  __   \
(_______)   |__|  \___)(_______/ |__|  \___)
`
}
