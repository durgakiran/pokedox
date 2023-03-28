package core

type CliCommnad struct {
	Name        string
	Description string
	Usage       string
	Command     func(args ...string) string
}

type Config struct {
	Commands map[string]CliCommnad
}

var CommandConfig Config = Config{
	Commands: make(map[string]CliCommnad),
}
