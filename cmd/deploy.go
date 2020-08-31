package cmd

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gush/parser"
	"gush/ssh"
	"log"
	"os"
	"os/exec"
	"strings"
)

var cfgFile string
var sshClient *ssh.SSHClient

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
		serverConfig, ok := parser.ServerConfig[server]
		if !ok {
			fmt.Println("Error: undefined server " + server)
			os.Exit(1)
		}

		taskConfig := getTaskConfig(task)

		color.Gray.Println("- connecting to the server...")
		sshClient = sshConn(serverConfig)
		defer sshClient.Close()

		runTask(task, taskConfig)
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
		fmt.Println("Error: error reading configuration file (" + err.Error() + ")")
		os.Exit(1)
	}
	if err := parser.ParseConfig(); err != nil {
		fmt.Println("Error: error parsing configuration file (" + err.Error() + ")")
		os.Exit(1)
	}

	color.Gray.Println("- using config file: " + viper.ConfigFileUsed())
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

		if action.Download.Local != "" && action.Download.Remote != "" {
			color.Green.Printf("[%s][download]\n", taskName)
			execDownload(action.Download)
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

func execDownload(download parser.DownloadAction) {
	color.Gray.Printf("- downloading file %s => %s \n", download.Remote, download.Local)
	err := sshClient.Download(download.Remote, download.Local)
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
