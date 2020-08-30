package parser

import (
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type UploadAction struct {
	Local  string `json:"local,omitempty" mapstructure:"local"`
	Remote string `json:"remote,omitempty" mapstructure:"remote"`
}

type DownloadAction struct {
	Local  string `json:"local,omitempty" mapstructure:"local"`
	Remote string `json:"remote,omitempty" mapstructure:"remote"`
}

type TaskAction struct {
	LocalShell  string         `json:"local_shell,omitempty" mapstructure:"local_shell"`
	RemoteShell string         `json:"remote_shell,omitempty" mapstructure:"remote_shell"`
	Upload      UploadAction   `json:"upload,omitempty"`
	Download    DownloadAction `json:"download,omitempty"`
	Run         []string       `json:"run,omitempty"`
}

type Task []TaskAction
type Tasks map[string]Task

func (t *Tasks) Parse() error {
	tasks := viper.Get("tasks")
	return mapstructure.Decode(tasks, t)
}
