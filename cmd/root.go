package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gush/util"
)

const VERSION = "1.0"

var showVersion = new(bool)

var rootCmd = &cobra.Command{
	Use:   "gush",
	Short: "Project deployment tool",
	Long:  `Project deployment tool.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		util.Fatalln(err.Error())
	}

	if *showVersion {
		fmt.Println("gush version " + VERSION)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(showVersion, "version", "v", false, "print the version")
}
