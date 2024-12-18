package cmd

import (
	"github.com/gookit/color"
	"github.com/jae-jae/gush/core/deploy"
	"github.com/jae-jae/gush/core/parser"
	"github.com/jae-jae/gush/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// deployCmd represents the deploy command
var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Project deployment",
	Long:  `Deploy the project through the configuration file.`,
	Run: func(cmd *cobra.Command, args []string) {
		task := "default"
		server := "default"
		if len(args) == 2 {
			task, server = args[0], args[1]
		} else if len(args) == 1 {
			task = args[0]
		}

		initConfig()

		deploy.RunTask(server, task)
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)

	deployCmd.Flags().StringVarP(&cfgFile, "config", "c", "./gush.yml", "config file")
}

func initConfig() {

	viper.SetConfigFile(cfgFile)

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		util.Fatalln("Error: error reading configuration file (" + err.Error() + ")")
	}
	if err := parser.ParseConfig(); err != nil {
		util.Fatalln("Error: error parsing configuration file (" + err.Error() + ")")
	}

	color.Gray.Println("- using config file: " + viper.ConfigFileUsed())
}
