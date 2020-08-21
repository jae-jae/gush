package parser

var TaskConfig Tasks

func ParseConfig() error {
	err := TaskConfig.Parse()
	if err != nil {
		return err
	}

	return nil
}
