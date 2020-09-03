package cmd

import (
	"github.com/spf13/cobra"
	"gush/util"
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
		util.Fatalln(err.Error())
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(showVersion, "version", "v", false, "print the version")
}
