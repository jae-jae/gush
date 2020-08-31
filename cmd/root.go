package cmd

import (
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"os"
)

const VERSION = "1.0"

var showVersion = new(bool)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gush",
	Short: "Project deployment tool",
	Long:  `Project deployment tool.`,
	//Run: func(cmd *cobra.Command, args []string) {
	//	if *showVersion {
	//		fmt.Println("gush version " + VERSION)
	//	}
	//},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		color.Red.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(showVersion, "version", "v", false, "print the version")
}
