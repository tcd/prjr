package prjr

import "strconv"

// VCS is an enum for different version control systems.
type VCS uint8

// VCS is an enum for different version control systems.
const (
	None VCS = iota // No vcs
	Git             // Git
	Hg              // Mercurial
	Svn             // Apache Subversion
)

func (vcs VCS) String() string {
	name := []string{"None", "Git", "Hg", "Svn"}
	i := uint8(vcs)
	switch {
	case i <= uint8(Svn):
		return name[i]
	default:
		return strconv.Itoa(int(i))
	}
}
