package parser

import (
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type ServerInfo struct {
	Host     string `json:"host,omitempty" mapstructure:"host"`
	User     string `json:"user,omitempty" mapstructure:"user"`
	Port     int    `json:"port,omitempty" mapstructure:"port"`
	Password string `json:"password,omitempty" mapstructure:"password"`
	SSHKey   string `json:"ssh_key,omitempty" mapstructure:"ssh_key"`
}

type Servers map[string]ServerInfo

func (s *Servers) Parse() error {
	servers := viper.Get("servers")
	return mapstructure.Decode(servers, s)
}
