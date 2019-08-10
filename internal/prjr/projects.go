package prjr

// Projects is a slice of Project with convenience methods.
type Projects struct {
	P []Project
}

// Add a Project to Projects.
func (pjs *Projects) Add(p ...Project) {
	for _, project := range p {
		pjs.P = append(pjs.P, project)
	}
}

// Save writes Projecs to a user's prjr.json file.
func (pjs Projects) Save() error {
	return writeProjectsToFile(getProjectsFilePath(), pjs.P)
}

// RemoveByRoot removes a Project identified by its Root.
func (pjs *Projects) RemoveByRoot(root string) {
	var newPjs Projects
	for _, project := range pjs.P {
		if project.Root != root {
			newPjs.Add(project)
		}
	}
	pjs.P = newPjs.P
}
