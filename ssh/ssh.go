package ssh

import (
	"fmt"
	"github.com/melbahja/goph"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
	"gush/parser"
	"strings"
	"time"
)

type SSHClient struct {
	*goph.Client
}

func ConnByConfig(config parser.ServerInfo) (*SSHClient, error) {
	var auth goph.Auth
	if config.Password != "" {
		auth = goph.Password(config.Password)
	} else if config.SSHKey != "" {
		auth = goph.Key(config.SSHKey, "")
	} else {
		password := askPass(fmt.Sprintf("%s@%s's password: ", config.User, config.Host))
		auth = goph.Password(password)
	}

	return Conn(config.User, config.Host, config.Port, auth)
}

func Conn(user string, addr string, port int, auth goph.Auth) (*SSHClient, error) {
	c := &goph.Client{
		Port: port,
		User: user,
		Addr: addr,
		Auth: auth,
	}

	err := goph.Conn(c, &ssh.ClientConfig{
		User:            c.User,
		Auth:            c.Auth,
		Timeout:         20 * time.Second,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})

	return &SSHClient{c}, err
}

func askPass(msg string) string {

	fmt.Print(msg)

	pass, err := terminal.ReadPassword(0)

	if err != nil {
		panic(err)
	}

	fmt.Println("")

	return strings.TrimSpace(string(pass))
}
