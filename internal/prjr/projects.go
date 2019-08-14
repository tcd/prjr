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

// Roots returns a slice of string with all Project Roots.
func (pjs Projects) Roots() []string {
	roots := make([]string, len(pjs.P))
	for i, project := range pjs.P {
		roots[i] = project.Root
	}
	return roots
}

// FindByRoot returns a Project identified by its Root,
// and a boolean which is false if no matching Project is found.
func (pjs Projects) FindByRoot(root string) (project Project, ok bool) {
	for _, pj := range pjs.P {
		if pj.Root == root {
			return pj, true
		}
	}
	return Project{}, false
}

// Names returns a slice of string with all Project Names.
func (pjs Projects) Names() []string {
	names := make([]string, len(pjs.P))
	for i, project := range pjs.P {
		names[i] = project.Name
	}
	return names
}

// FindByName returns a Project identified by its Name,
// and a boolean which is false if no matching Project is found.
func (pjs Projects) FindByName(name string) (project Project, ok bool) {
	for _, pj := range pjs.P {
		if pj.Root == name {
			return pj, true
		}
	}
	return Project{}, false
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

// GetLocalProjects returns the contents of a user's prjr.json file
// as a Products struct.
func GetLocalProjects() (Projects, error) {
	var projects Projects
	existingProjects, err := GetProjects()
	if err != nil {
		return projects, err
	}
	projects.Add(existingProjects...)
	return projects, nil
}
