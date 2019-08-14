package stat

import (
	"strconv"
	"strings"
)

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

// AheadString returns the Ahead field as a string.
func (gs GitStatus) AheadString() string {
	return strconv.Itoa(gs.Ahead)
}

// BehindString returns the Behind field as a string.
func (gs GitStatus) BehindString() string {
	return strconv.Itoa(gs.Ahead)
}

// StashString returns the Stashes field as a string.
func (gs GitStatus) StashString() string {
	return strconv.Itoa(gs.Stashed)
}

func (gs GitStatus) String() string {
	var s []string // This isn't too efficient but it will, at most, reallocate s 12 times.

	if gs.BranchName != "" {
		s = append(s, "Branch: "+gs.BranchName)
	}
	if gs.Ahead > 0 {
		s = append(s, "Ahead by "+gs.AheadString())
	}
	if gs.Behind > 0 {
		s = append(s, "Behind by "+gs.BehindString())
	}
	if gs.Modified {
		s = append(s, "Modified")
	}
	if gs.Added {
		s = append(s, "Added")
	}
	if gs.Deleted {
		s = append(s, "Deleted")
	}
	if gs.UpdatedUnmerged {
		s = append(s, "Unmerged")
	}
	if gs.Untracked {
		s = append(s, "Untracked")
	}
	if gs.Staged {
		s = append(s, "Staged")
	}
	if gs.Stashed > 0 {
		if gs.Stashed == 1 {
			s = append(s, "1 Stash")
		} else {
			s = append(s, string(gs.StashString())+" Stashes")
		}
	}
	return strings.Join(s, ", ")
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
	if gs.Ahead > 0 {
		sb.WriteString(Icons.Ahead)
		sb.WriteString(string(gs.Ahead))
	}
	if gs.Behind > 0 {
		sb.WriteString(Icons.Behind)
		sb.WriteString(string(gs.Behind))
	}
	if gs.Modified {
		sb.WriteString(Icons.Modified)
	}
	if gs.Added {
		sb.WriteString(Icons.Added)
	}
	if gs.Deleted {
		sb.WriteString(Icons.Deleted)
	}
	if gs.UpdatedUnmerged {
		sb.WriteString(Icons.Unmerged)
	}
	if gs.Untracked {
		sb.WriteString(Icons.Untracked)
	}
	if gs.Staged {
		sb.WriteString(Icons.Staged)
	}
	if gs.Stashed > 0 {
		sb.WriteString(Icons.Stashed)
	}

	return sb.String()
}
