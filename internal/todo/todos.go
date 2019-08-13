package todo

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

// Todos recursively searches files for TODO comments
// and returns all TODOs.
func Todos(path string) []Todo {
	var todos []Todo

	filePaths := ListFilesRecursive(path)
	files := make([][]byte, len(filePaths))

	for i, file := range filePaths {
		bytes, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println("error reading file:", err)
		}
		files[i] = bytes
	}

	pattern := searchString(todoStrings)
	re := regexp.MustCompile(pattern)

	for i, file := range files {
		matches := re.FindAllStringSubmatch(string(file), -1)
		if matches != nil {
			for _, match := range matches {
				todo := Todo{
					File:    filePaths[i],
					Type:    match[1],
					Author:  match[3],
					Content: match[4],
				}
				todos = append(todos, todo)
			}
		}
	}

	return todos
}

// Returns a Regular Expression that can be used to identify TODO comments.
func searchString(todos []string) string {
	return fmt.Sprintf(`(?P<type>%s)(\((?P<author>\w+)\))?:\s(?P<content>.*)`, strings.Join(todos, "|"))
}
