package deploy

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/gookit/color"
	"github.com/jae-jae/gush/core/parser"
	"github.com/jae-jae/gush/util"
)

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
		util.Fatalln(err.Error())
	}
}

func execDownload(download parser.DownloadAction) {
	color.Gray.Printf("- downloading file %s => %s \n", download.Remote, download.Local)
	err := sshClient.Download(download.Remote, download.Local)
	if err != nil {
		util.Fatalln(err.Error())
	}
}

func execRemoteShell(shell string) {
	shell = echoCommand(shell)
	out, err := sshClient.Run(shell)
	if err != nil {
		util.Fatalln(string(out))
	}
	color.Cyan.Println(string(out))
}

func execLocalShell(shell string) {
	shell = echoCommand(shell)
	cmd := exec.Command("/bin/sh", "-c", shell)
	b, err := cmd.CombinedOutput()
	if err != nil {
		util.Fatalln(string(b))
	}
	color.Blue.Println(string(b))
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
