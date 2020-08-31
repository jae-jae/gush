package cmd

import (
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"io/ioutil"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Generate configuration file",
	Long:  `Generate configuration file.`,
	Run: func(cmd *cobra.Command, args []string) {
		tpl := configTemplate()
		err := ioutil.WriteFile("./gushx.yml", []byte(tpl), 0777)
		if err != nil {
			color.Red.Println(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

}

func configTemplate() string {
	return `servers:
  default:
    host: site.com
    user: root
    port: 22
    password: ""
    ssh_key: ""

tasks:
  default:
    - local_shell: |
        cd /www/
        git push origin master
    - upload:
        local: /path/to/build.sh
        remote: /wwwroot/build.sh
    - remote_shell: |
        cd /wwwroot/
        git pull origin master
        yarn build
    - run:
        - task_2
        - task_3
  task_2:
    - local_shell: |
        echo "task 1"
  task_3:
    - local_shell: |
        echo "task 2"
    - download:
        remote: /wwwroot/build.sh
        local: /www/build.sh
    - remote_shell: |
        rm /wwwroot/build.sh
        ls -la /wwwroot/
`
}
