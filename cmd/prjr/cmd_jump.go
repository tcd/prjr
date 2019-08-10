package main

import (
	"fmt"
	"os"

	"github.com/tcd/prjr/internal/prjr"
)

func jumpCmd(pjs prjr.Projects) {
	home, _ := os.UserHomeDir()
	fmt.Println("jumpto:" + home)
}
