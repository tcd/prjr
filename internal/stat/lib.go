package stat

import (
	"strconv"
	"strings"
)

// GetGitStatus returns information about a local Git repository.
func GetGitStatus(path string) (GitStatus, error) {
	cmdOutput, err := gitStatusCmd(path)
	if err != nil {
		return GitStatus{}, err
	}

	branchInfo, fileInfo := parseGitStatusPorcelainV2(cmdOutput)
	stashCount := getStashes(path)

	return GitStatus{
		GitBranchInfo: branchInfo,
		GitFileInfo:   fileInfo,
		Stashed:       stashCount,
	}, nil
}

func gitStatusCmd(path string) (string, error) {
	cmdName := "git"
	porcelainArgs := []string{"-C", path, "status", "--porcelain=v2", "-b"} // TODO: Ignore submodules?
	output, err := runCmd(cmdName, porcelainArgs...)
	if err != nil {
		return "", err
	}
	return output, nil
}

// Parses the output of `git status --porcelain=v2 -b`
func parseGitStatusPorcelainV2(output string) (GitBranchInfo, GitFileInfo) {
	lines := strings.Split(output, "\n")
	var branchLines, fileLines []string
	for _, line := range lines {
		if len(line) > 0 {
			if []rune(line)[0] == '#' {
				branchLines = append(branchLines, line)
			} else if []rune(line)[0] != ' ' {
				fileLines = append(fileLines, line)
			}
		}
	}
	branchInfo := parseBranchInfo(branchLines)
	fileInfo := parseFileInfo(fileLines)
	return branchInfo, fileInfo
}

func parseBranchInfo(lines []string) GitBranchInfo {
	var branchInfo GitBranchInfo
	for _, line := range lines {
		// current local branch
		if strings.Contains(line, "branch.head") {
			fields := strings.Fields(line)
			branch := fields[len(fields)-1]
			branchInfo.BranchName = branch
		}
		// current upstream branch
		if strings.Contains(line, "branch.upstream") {
			fields := strings.Fields(line)
			branch := fields[len(fields)-1]
			branchInfo.BranchUpstreamName = branch
		}
		// ahead/behind
		if strings.Contains(line, "branch.ab") {
			fields := strings.Fields(line)
			if len(fields[2]) >= 2 {
				ahead := fields[2][1:]
				aheadCount, _ := strconv.Atoi(strings.TrimSpace(ahead))
				branchInfo.Ahead = aheadCount
			}
			if len(fields[3]) >= 2 {
				behind := fields[3][1:]
				behindCount, _ := strconv.Atoi(strings.TrimSpace(behind))
				branchInfo.Behind = behindCount
			}
		}
	}
	return branchInfo
}

func parseFileInfo(lines []string) GitFileInfo {
	var fileInfo GitFileInfo

	for _, line := range lines {
		fields := strings.Fields(line)

		// Ordinary changed entries
		if fields[0] == "1" {
			info := fields[1]
			switch info {
			// Added
			case ".A":
				fileInfo.Added = true
			case "A.":
				fileInfo.Added = true
				fileInfo.Staged = true
			// Deleted
			case ".D":
				fileInfo.Deleted = true
			case "D.":
				fileInfo.Deleted = true
				fileInfo.Staged = true
			// Modified
			case ".M":
				fileInfo.Modified = true
			case "M.":
				fileInfo.Modified = true
				fileInfo.Staged = true
			}
			continue
		}
		// Renamed or copied entries
		if fields[0] == "2" {
			info := fields[1]
			switch info {
			// Renamed
			case ".R":
				fileInfo.Renamed = true
			case "R.":
				fileInfo.Renamed = true
				fileInfo.Staged = true
			// Copied
			case ".C":
				fileInfo.Copied = true
			case "C.":
				fileInfo.Copied = true
				fileInfo.Staged = true
			}
			continue
		}
		// Unmerged entries
		if fields[0] == "3" {
			fileInfo.UpdatedUnmerged = true
			continue
		}
		// Untracked items
		if fields[0] == "?" {
			fileInfo.Untracked = true
			continue
		}
		// Ignored items
		if fields[0] == "!" {
			continue
		}
	}

	return fileInfo
}

func getStashes(path string) int {
	cmdName := "git"
	stashCountArgs := []string{"-C", path, "rev-list", "--walk-reflogs", "--count", "refs/stash"}
	output, err := runCmd(cmdName, stashCountArgs...)
	if err != nil {
		// TODO: Handle individual errors.
		return 0
	}
	parsedCount, _ := strconv.Atoi(strings.TrimSpace(output))
	return parsedCount
}
