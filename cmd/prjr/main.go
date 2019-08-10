package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tcd/prjr/internal/prjr"
)

func init() {
	flag.Usage = usage
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
		case "list":
			listCmd(projects)
		case "remove":
			rmCmd(projects)
		default:
			flag.Usage()
		}
	} else {
		flag.Usage()
	}
}

func usage() {
	fmt.Println(titleString())
	fmt.Printf("Usage: %s <command> [options]\n", os.Args[0])
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
