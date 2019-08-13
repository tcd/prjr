package todo

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

// TodosCount recursively searches files for TODO comments
// and returns the number of TODOs found.
// Should be faster than Todos().
func TodosCount(path string) int {
	var count int

	filePaths := ListFilesRecursive(path)
	files := make([][]byte, len(filePaths))

	for i, file := range filePaths {
		bytes, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println("error reading file:", err)
		}
		files[i] = bytes
	}

	pattern := simpleSearchString(todoStrings)
	re := regexp.MustCompile(pattern)

	for _, file := range files {
		matches := re.FindAllStringSubmatch(string(file), -1)
		if matches != nil {
			for range matches {
				count++
			}
		}
	}

	return count
}

// Returns a Regular Expression that can be used to identify TODO comments.
func simpleSearchString(todos []string) string {
	return fmt.Sprintf(`(%s)(\((\w+)\))?:\s(.*)`, strings.Join(todos, "|"))
}
