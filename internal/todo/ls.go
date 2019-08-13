package todo

import (
	"fmt"
	"os"
	"path/filepath"
)

// ListFilesRecursive returns the full path of all files within a folder.
func ListFilesRecursive(targetPath string) []string {
	var paths []string
	os.Chdir(targetPath)
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && contains(ignoredFolders, info.Name()) {
			return filepath.SkipDir
		}
		if !info.IsDir() && !contains(ignoredFiles, info.Name()) {
			paths = append(paths, filepath.Join(targetPath, path))
		}
		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", targetPath, err)
		return paths
	}
	return paths
}

// check if a slice of string contains a specific string.
func contains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}
