package parser

var TaskConfig Tasks
var ServerConfig Servers

func ParseConfig() error {
	err := TaskConfig.Parse()
	if err != nil {
		return err
	}

	err = ServerConfig.Parse()
	if err != nil {
		return err
	}

	return nil
}
