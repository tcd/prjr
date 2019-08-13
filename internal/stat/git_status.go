package stat

// GitBranchInfo contains info about a local Git repository's current branch.
type GitBranchInfo struct {
	BranchName         string
	BranchUpstreamName string
	Ahead              int
	Behind             int
}

// GitFileInfo contains information about files in a Git repository.
type GitFileInfo struct {
	Modified        bool
	Added           bool
	Deleted         bool
	Renamed         bool
	Copied          bool
	UpdatedUnmerged bool
	Untracked       bool
	Staged          bool
}

// GitStatus models the information about a git reposiory's state.
type GitStatus struct {
	GitBranchInfo
	// Branch string
	// Ahead  int
	// Behind int

	GitFileInfo
	// Modified        bool
	// Added           bool
	// Deleted         bool
	// Renamed         bool
	// Copied          bool
	// UpdatedUnmerged bool
	// Untracked       bool
	// Staged          bool

	Stashed int
}

// Diverged is true if a branch is both ahead of & behind upstream.
func (gs GitStatus) Diverged() bool {
	if gs.Ahead > 0 && gs.Behind > 0 {
		return true
	}
	return false
}
