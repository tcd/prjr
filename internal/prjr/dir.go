package prjr

import (
	"io/ioutil"
	"path/filepath"
)

// Dir contains information about and logic for dealing with a directory
type Dir struct {
	Path    string
	Files   []string
	Folders []string
}

// Has checks to see if Dir contains a file or folder with a given name.
func (d Dir) Has(name string) bool {
	for _, file := range d.Files {
		if file == name {
			return true
		}
	}
	for _, folder := range d.Folders {
		if folder == name {
			return true
		}
	}
	return false
}

// HasFile checks to see if Dir contains a file a given name.
func (d Dir) HasFile(fileName string) bool {
	for _, file := range d.Files {
		if file == fileName {
			return true
		}
	}
	return false
}

// HasFolder checks to see if Dir contains a folder a given name.
func (d Dir) HasFolder(folderName string) bool {
	for _, folder := range d.Folders {
		if folder == folderName {
			return true
		}
	}
	return false
}

func newDir(path string) (Dir, error) {
	var dir Dir

	absPath, err := filepath.Abs(path)
	if err != nil {
		return dir, err
	}
	dir.Path = absPath

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return dir, err
	}
	for _, f := range files {
		if f.IsDir() {
			dir.Folders = append(dir.Folders, f.Name())
		} else {
			dir.Files = append(dir.Files, f.Name())
		}
	}

	return dir, nil
}
