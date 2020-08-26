package parser

import (
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type TaskAction struct {
	LocalShell  string            `json:"local_shell,omitempty" mapstructure:"local_shell"`
	RemoteShell string            `json:"remote_shell,omitempty" mapstructure:"remote_shell"`
	Upload      map[string]string `json:"upload,omitempty"`
	Download    map[string]string `json:"download,omitempty"`
	Run         []string          `json:"run,omitempty"`
}

type Task []TaskAction
type Tasks map[string]Task

func (t *Tasks) Parse() error {
	tasks := viper.Get("tasks")
	return mapstructure.Decode(tasks, t)
}
