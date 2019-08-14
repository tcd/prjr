package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:       "completion",
	Short:     "Generate shell completion files for prjr",
	Long:      `Generate shell completion files for prjr`,
	Hidden:    true,
	Args:      cobra.ExactArgs(1),
	ValidArgs: []string{"bash", "zsh"},
	Run: func(cmd *cobra.Command, args []string) {
		if args[0] == "bash" {
			rootCmd.GenBashCompletion(os.Stdout)
		}
		if args[0] == "zsh" {
			rootCmd.GenZshCompletion(os.Stdout)
		}
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
