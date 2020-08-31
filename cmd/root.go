package cmd

import (
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"os"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gush",
	Short: "Project deployment tool",
	Long:  `Project deployment tool.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		color.Red.Println(err)
		os.Exit(1)
	}
}

func init() {

}
