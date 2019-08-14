package prjr

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func newProjectFromDirectory(path string) (Project, error) {
	var pj Project

	pjContents, err := readDir(path)
	if err != nil {
		return pj, err
	}
	for _, fileName := range pjContents {
		if fileName == ".git" {
			pj.VCS = true
		}
	}
	pj.Root = path
	pj.Name = filepath.Base(path)
	return pj, nil
}

func readProjectsFromFile(path string) ([]Project, error) {
	var projects []Project
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return projects, err
	}
	err = json.Unmarshal(bytes, &projects)
	if err != nil {
		return projects, err
	}
	return projects, nil
}

func writeProjectsToFile(path string, projects []Project) error {
	bytes, err := json.MarshalIndent(projects, "", "  ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, bytes, os.FileMode(0644))
	if err != nil {
		return err
	}
	return nil
}

func getProjectsFilePath() string {
	home := os.Getenv("HOME")
	dataHome := os.Getenv("XDG_DATA_HOME")
	prjrDir := os.Getenv("PRJR_DIR")

	if prjrDir == "" {
		if home == "" {
			home = "~"
		}
		if dataHome == "" {
			dataHome = filepath.Join(home, ".local", "share")
		}
		prjrDir = filepath.Join(dataHome, "prjr")
	}

	if _, err := os.Stat(prjrDir); os.IsNotExist(err) {
		err = os.MkdirAll(prjrDir, 0777)
		if err != nil {
			fmt.Printf("Error creating projects directory: %s\n", err)
		}
	}

	prjrFile := filepath.Join(prjrDir, "prjr.json")
	if _, err := os.Stat(prjrFile); os.IsNotExist(err) {
		err = writeProjectsToFile(prjrFile, []Project{})
		if err != nil {
			fmt.Printf("Error creating projects file: %s\n", err)
		}
	}

	return filepath.Join(prjrDir, "prjr.json")
}

// return a string slice with the names of all files & folders in a directory
func readDir(path string) ([]string, error) {
	var fileNames []string
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return fileNames, err
	}
	for _, f := range files {
		// if !f.IsDir() {
		fileNames = append(fileNames, f.Name())
		// }
	}
	return fileNames, nil
}
