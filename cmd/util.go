package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var utilCmd = &cobra.Command{
	Use:       "util [bash|zsh|man]",
	Short:     "Generate completion files or manual pages for prjr",
	Hidden:    true,
	Args:      cobra.ExactArgs(1),
	ValidArgs: []string{"bash", "zsh", "man"},
	Run: func(cmd *cobra.Command, args []string) {
		if args[0] == "bash" {
			rootCmd.GenBashCompletion(os.Stdout)
			os.Exit(0)
		}
		if args[0] == "zsh" {
			rootCmd.GenZshCompletion(os.Stdout)
			os.Exit(0)
		}
	},
}

func init() {
	rootCmd.AddCommand(utilCmd)
}

func genMan(cmd *cobra.Command, path string) {
	if filepath.IsAbs(path) {
		header := &doc.GenManHeader{
			Title:   "PRJR",
			Section: "1",
			Source:  "",
		}
		doc.GenManTree(cmd, header, path)
		os.Exit(0)
	} else {
		fmt.Fprintln(os.Stderr, "Please provide an absolute path")
		os.Exit(1)
	}
}
