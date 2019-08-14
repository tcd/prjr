package cmd

// check if a slice of string contains a specific string.
func contains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}
