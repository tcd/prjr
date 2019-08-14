package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var utilCmd = &cobra.Command{
	Use:    "util",
	Short:  "Generate completion files or manual pages for prjr",
	Hidden: true,
	Args:   cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		zsh, _ := cmd.Flags().GetBool("zsh")
		if zsh {
			rootCmd.GenZshCompletion(os.Stdout)
			os.Exit(0)
		}

		bash, _ := cmd.Flags().GetBool("bash")
		if bash {
			rootCmd.GenBashCompletion(os.Stdout)
			os.Exit(0)
		}

	},
}

func init() {
	rootCmd.AddCommand(utilCmd)
	utilCmd.Flags().Bool("bash", false, "Print prjr bash completions to stdout")
	utilCmd.Flags().Bool("zsh", false, "Print prjr zsh completions to stdout")
	utilCmd.Flags().String("man", "", "Path to a directory prjr manpages should be added to")
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
