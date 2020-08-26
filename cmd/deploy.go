/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gush/parser"
	"gush/ssh"
	"os"
)

var sshClient *ssh.SSHClient

// deployCmd represents the deploy command
var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		task := "default"
		server := "default"
		if len(args) == 2 {
			task, server = args[0], args[1]
		} else if len(args) == 1 {
			task = args[0]
		}

		serverConfig, ok := parser.ServerConfig[server]
		if !ok {
			fmt.Println("Error: undefined server " + server)
			os.Exit(1)
		}

		taskConfig, ok := parser.TaskConfig[task]
		if !ok {
			fmt.Println("Error: undefined task " + task)
			os.Exit(1)
		}

		//fmt.Printf("%#v\n", serverConfig)
		//fmt.Printf("%#v\n", taskConfig)

		fmt.Println("Connecting to the server...")
		sshClient = sshConn(serverConfig)
		//b, _ := client.Run("ls -la")
		//fmt.Println(string(b))
		runTask(taskConfig)
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deployCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deployCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func sshConn(config parser.ServerInfo) *ssh.SSHClient {
	client, err := ssh.ConnByConfig(config)
	if err != nil {
		fmt.Println("Error: error connecting to server （" + err.Error() + ")")
		os.Exit(1)
	}

	return client
}

func runTask(task parser.Task) {
	for _, action := range task {
		if action.LocalShell != "" {
			execLocalShell(action.LocalShell)
		}

		if action.RemoteShell != "" {
			execRemoteShell(action.RemoteShell)
		}

		if len(action.Upload) != 0 {
			execUpload(action.Upload)
		}

		if len(action.Run) != 0 {
			execRun(action.Run)
		}
	}
}

func execRun(tasks []string) {

}

func execUpload(upload map[string]string) {

}

func execRemoteShell(shell string) {

}

func execLocalShell(shell string) {

}
