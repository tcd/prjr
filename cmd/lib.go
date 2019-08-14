package cmd

import "strings"

// check if a slice of string contains a specific string.
func contains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}

// converts all strings in a slice to lowercase.
func toLower(slice []string) []string {
	newSlice := make([]string, len(slice))
	for i, s := range slice {
		newSlice[i] = strings.ToLower(s)
	}
	return newSlice
}

// Returns a fancy ascii title.
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
