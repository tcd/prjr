package prjr

import (
	"fmt"
	"os"
	"time"

	"github.com/tcd/prjr/internal/stat"
)

// Project is the main datastructure of the prjr application.
type Project struct {
	Root      string    `json:"root"`      // Path to the the Project's root directory on the local machine. TODO: expand env variables in root? Need constructor function?
	Name      string    `json:"name"`      // Name for describing the Project
	Timestamp time.Time `json:"timestamp"` // Timestamp used to identify the project
	VCS       bool      `json:"vcs"`       // Whether or not the Project is under version control
	Favorite  bool      `json:"favorite"`  // Favorite projects are listed first
	Fork      bool      `json:"fork"`      // True if the project is forked from another repository
}

// NewProjectHere returns a new Project for the current directory.
func NewProjectHere() (Project, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return Project{}, err
	}
	return newProjectFromDirectory(cwd)
}

// GetProjects fetches Projects from a user's prjr.json file.
func GetProjects() ([]Project, error) {
	return readProjectsFromFile(getProjectsFilePath())
}

// SaveProjects writes a slice of Project to a user's prjr.json file.
func SaveProjects(projects []Project) error {
	return writeProjectsToFile(getProjectsFilePath(), projects)
}

// GitStatus returns information about a local Git repository.
func (p Project) GitStatus() (stat.GitStatus, error) {
	if p.VCS == false {
		return stat.GitStatus{}, fmt.Errorf("error: Project is not a git repository")
	}
	stats, err := stat.GetGitStatus(p.Root)
	if err != nil {
		return stat.GitStatus{}, err
	}
	return stats, nil
}
