package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "prjr",
	Short: "A command line project manager",
	Long:  titleString(),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}

		outputPath, _ := cmd.Flags().GetString("manpages")
		if outputPath != "" {
			genMan(cmd, outputPath)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main().
// It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// cobra.OnInitialize(initConfig)

	// Cobra supports persistent flags, which, if defined here, will be global for your application.
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.prjr.yaml)")

	rootCmd.Flags().BoolP("version", "v", false, "Output version information about prjr")
	rootCmd.Flags().String("manpages", "", "Output prjr man pages to a given directory")
	rootCmd.Flags().MarkHidden("manpages")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".prjr" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".prjr")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
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
