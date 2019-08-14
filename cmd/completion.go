package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:       "completion [bash|zsh]",
	Short:     "Generate shell completion files for prjr",
	Hidden:    true,
	Args:      cobra.ExactArgs(1),
	ValidArgs: []string{"bash", "zsh"},
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
	rootCmd.AddCommand(completionCmd)
}
