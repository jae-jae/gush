package cmd

import (
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gush/core/deploy"
	"gush/core/parser"
	"gush/util"
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

	rootCmd.Flags().StringVar(&cfgFile, "config", "", "config file (default is ./gush.yaml)")
}

func initConfig() {
	if cfgFile == "" {
		cfgFile = "./gush.yml"
	}

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
