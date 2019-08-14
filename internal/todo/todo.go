package todo

import "strings"

// Todo comment in a file.
type Todo struct {
	File    string // Full path to the file the TODO is from.
	RelPath string // Path to the file containing the todo relative to the project root.
	Type    string // Type of TODO; ex: TODO, FIXME, etc.
	Author  string // Name of the person who left the TODO, if present.
	Content string // The comment following the TODO keywork.
}

func (t Todo) String() string {
	var sb strings.Builder

	sb.WriteString(t.Type)
	if t.Author != "" {
		sb.WriteString("(" + t.Author + ")")
	}
	sb.WriteString(": ")
	sb.WriteString(t.Content)
	sb.WriteString("\t(" + t.File + ")")

	return sb.String()
}

// todoStrings are different TODOs to search for.
var todoStrings = []string{
	"TODO",
	"FIXME",
	"BUG",
	"NOTE",
}

// don't search these folders for TODOs
var ignoredFolders = []string{
	".git",
	"node_modules",
	"bower_components",
	"vendor",
	"__pycache__",
	".mypy_cache",
	".sass-cache",
	"dist",
	"undo",
}

// don't search these files for TODOs
var ignoredFiles = []string{
	".DS_Store",
	"tags",
}

// don't search files with these extensions for TODOs
var ignoredExtensions = []string{
	// compressed files
	".tgz",
	".zip",
	// compiled/built files
	".beam",
	".dll",
	".dylib",
	".exe",
	".obj",
	".out",
	".so",
	// eBook build output
	".epub",
	".mobi",
	".pdf",
	// misc
	".png",
	".rs.bk",
}
