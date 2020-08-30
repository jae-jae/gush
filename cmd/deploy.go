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
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"gush/parser"
	"gush/ssh"
	"log"
	"os"
	"os/exec"
	"strings"
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

		taskConfig := getTaskConfig(task)

		//fmt.Printf("%#v\n", serverConfig)
		fmt.Printf("%#v\n", taskConfig)

		color.Gray.Println("- connecting to the server...")
		sshClient = sshConn(serverConfig)
		//b, _ := client.Run("ls -la")
		//fmt.Println(string(b))
		runTask(task, taskConfig)
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

func getTaskConfig(task string) parser.Task {
	taskConfig, ok := parser.TaskConfig[task]
	if !ok {
		fmt.Println("Error: undefined task " + task)
		os.Exit(1)
	}

	return taskConfig
}

func sshConn(config parser.ServerInfo) *ssh.SSHClient {
	client, err := ssh.ConnByConfig(config)
	if err != nil {
		fmt.Println("Error: error connecting to server （" + err.Error() + ")")
		os.Exit(1)
	}

	return client
}

func runTask(taskName string, task parser.Task) {
	color.Gray.Println("- run task " + taskName)
	for _, action := range task {
		if action.LocalShell != "" {
			color.Green.Printf("[%s][local_shell]\n", taskName)
			execLocalShell(action.LocalShell)
		}

		if action.RemoteShell != "" {
			color.Green.Printf("[%s][remote_shell]\n", taskName)
			execRemoteShell(action.RemoteShell)
		}

		if action.Upload.Local != "" && action.Upload.Remote != "" {
			color.Green.Printf("[%s][upload]\n", taskName)
			execUpload(action.Upload)
		}

		if len(action.Run) != 0 {
			color.Green.Printf("[%s][run]\n", taskName)
			execRun(action.Run)
		}
	}
}

func execRun(tasks []string) {
	for _, task := range tasks {
		taskConfig := getTaskConfig(task)
		runTask(task, taskConfig)
	}
}

func execUpload(upload parser.UploadAction) {
	color.Gray.Printf("- uploading file %s => %s \n", upload.Local, upload.Remote)
	err := sshClient.Upload(upload.Local, upload.Remote)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func execRemoteShell(shell string) {
	shell = echoCommand(shell)
	out, err := sshClient.Run(shell)
	if err != nil {
		log.Fatalln(err.Error())
	}
	color.Cyan.Println(string(out))
}

func execLocalShell(shell string) {
	shell = echoCommand(shell)
	cmd := exec.Command("/bin/sh", "-c", shell)
	b, err := cmd.Output()
	if err != nil {
		log.Fatalln(err.Error())
	}
	color.Black.Println(string(b))
}

// 回显命令
func echoCommand(shell string) string {
	commands := strings.Split(shell, "\n")
	tmp := []string{}
	for _, command := range commands {
		if command != "" {
			tmp = append(tmp, fmt.Sprintf("echo \"> %s\"", command), command)
		}
	}

	return strings.Join(tmp, ";")
}
