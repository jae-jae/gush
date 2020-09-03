package deploy

import (
	"github.com/gookit/color"
	"gush/core/parser"
	"gush/core/ssh"
	"gush/util"
)

var sshClient *ssh.SSHClient

func RunTask(server string, task string) {
	serverConfig := getServerConfig(server)
	taskConfig := getTaskConfig(task)
	sshClient = sshConn(serverConfig)
	defer sshClient.Close()
	runTask(task, taskConfig)
}

func sshConn(config parser.ServerInfo) *ssh.SSHClient {
	color.Gray.Println("- connecting to the server...")
	client, err := ssh.ConnByConfig(config)
	if err != nil {
		util.Fatalln("Error: error connecting to server （" + err.Error() + ")")
	}

	return client
}

func getTaskConfig(task string) parser.Task {
	taskConfig, ok := parser.TaskConfig[task]
	if !ok {
		util.Fatalln("Error: undefined task " + task)
	}

	return taskConfig
}

func getServerConfig(server string) parser.ServerInfo {
	serverConfig, ok := parser.ServerConfig[server]
	if !ok {
		util.Fatalln("Error: undefined server " + server)
	}

	return serverConfig
}
