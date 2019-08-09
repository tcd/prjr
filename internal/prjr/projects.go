package prjr

// Projects is a slice of Project with convenience methods.
type Projects []Project

// Add a Project to Projects.
func (pjs *Projects) Add(p ...Project) {
	for _, project := range p {
		*pjs = append(*pjs, project)
	}
}
