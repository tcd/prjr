package stat

import "strings"

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

// GitStatusIconSet is used to hold icons for use in printing.
type GitStatusIconSet struct {
	Untracked string
	Added     string
	Modified  string
	Staged    string
	Deleted   string
	Stashed   string
	Unmerged  string
	Ahead     string
	Behind    string
	Diverged  string
}

// Icons are the default icons for short GitStatus printing.
var Icons = GitStatusIconSet{
	Untracked: "?",
	Added:     "+",
	Modified:  "!",
	Staged:    "§",
	Deleted:   "✘",
	Stashed:   "$",
	Unmerged:  "=",
	Ahead:     "⇡",
	Behind:    "⇣",
	Diverged:  "⇕",
}

// IconString returns a string of unicode icons representing GitStatus.
func (gs GitStatus) IconString() string {
	var sb strings.Builder

	// Branch string
	if gs.BranchName != "" {
		sb.WriteString(gs.BranchName + " ")
	}
	// Ahead  int
	if gs.Ahead > 0 {
		sb.WriteString(Icons.Ahead)
		sb.WriteString(string(gs.Ahead))
	}
	// Behind int
	if gs.Behind > 0 {
		sb.WriteString(Icons.Behind)
		sb.WriteString(string(gs.Behind))
	}
	// Modified bool
	if gs.Modified {
		sb.WriteString(Icons.Modified)
	}
	// Added bool
	if gs.Added {
		sb.WriteString(Icons.Added)
	}
	// Deleted bool
	if gs.Deleted {
		sb.WriteString(Icons.Deleted)
	}
	// Renamed bool
	// Copied bool
	// UpdatedUnmerged bool
	if gs.UpdatedUnmerged {
		sb.WriteString(Icons.Unmerged)
	}
	// Untracked bool
	if gs.Untracked {
		sb.WriteString(Icons.Untracked)
	}
	// Staged bool
	if gs.Staged {
		sb.WriteString(Icons.Staged)
	}
	// Stashed int
	if gs.Stashed > 0 {
		sb.WriteString(Icons.Stashed)
	}

	return sb.String()
}
